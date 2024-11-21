package services

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	c "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	pkgUtl "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type GrantType string

const (
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeRefreshToken      GrantType = "refresh_token"
)

type FetchOAuthTokensParams struct {
	GrantType              GrantType
	ClientId               string
	CodeVerifier           string
	AuthCodeOrRefreshToken string
}

/* FetchOAuthTokens return (access_token, refresh_token, error)
 */
func FetchOAuthTokens(p FetchOAuthTokensParams) (string, string, error) {
	// form fields
	values := url.Values{}
	values.Add("client_id", p.ClientId)
	values.Add("code_verifier", p.CodeVerifier)
	values.Add("grant_type", string(p.GrantType))
	if p.GrantType == GrantTypeAuthorizationCode {
		values.Add("code", p.AuthCodeOrRefreshToken)
	} else {
		values.Add("refresh_token", p.AuthCodeOrRefreshToken)
	}

	res, err := http.PostForm("https://myanimelist.net/v1/oauth2/token", values)
	if err != nil {
		return "", "", err
	}

	var data struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Println("error decoding body into valid json")
		return "", "", err
	}
	defer res.Body.Close()

	return data.AccessToken, data.RefreshToken, nil
}

func UpdateAccessToken() error {
	log.Printf("fetching new access_token from mal api...")
	accessToken, refreshToken, err := FetchOAuthTokens(FetchOAuthTokensParams{
		GrantType:              GrantTypeRefreshToken,
		ClientId:               c.C.MalClientId,
		CodeVerifier:           c.C.MalCodeVerifier,
		AuthCodeOrRefreshToken: pkgUtl.ReadConfigFile("refresh_token"),
	})

	if err != nil {
		return err
	}

	//log.Println("access_token: ", accessToken)
	//log.Println("refresh_token: ", accessToken)

	if err := pkgUtl.WriteConfigFile("access_token", accessToken); err != nil {
		log.Println("error writing access_token to config file", err)
	}
	if err := pkgUtl.WriteConfigFile("refresh_token", refreshToken); err != nil {
		log.Println("error writing refresh_token to config file", err)
	}

	log.Println("updated access_token and refresh_token in config file")

	return nil
}

func UpdateRefreshToken() error {
	if c.C.MalAuthCode == "" {
		c.C.MalAuthCode = pkgUtl.ReadConfigFile("auth_code")
		if c.C.MalAuthCode == "" {
			// user has not given permissions to access there data.
			log.Println("auth_code is missing in config file")
			return errors.New("failed to update refresh_token")
		}
	}

	log.Printf("fetching new refresh_token from mal api...")
	accessToken, refreshToken, err := FetchOAuthTokens(FetchOAuthTokensParams{
		GrantType:              GrantTypeAuthorizationCode,
		ClientId:               c.C.MalClientId,
		CodeVerifier:           c.C.MalCodeVerifier,
		AuthCodeOrRefreshToken: c.C.MalAuthCode,
	})

	if err != nil {
		return err
	}

	//log.Println("access_token: ", accessToken)
	//log.Println("refresh_token: ", accessToken)

	if err := pkgUtl.WriteConfigFile("access_token", accessToken); err != nil {
		log.Println("error writing access_token to config file", err)
	}
	if err := pkgUtl.WriteConfigFile("refresh_token", refreshToken); err != nil {
		log.Println("error writing refresh_token to config file", err)
	}

	log.Println("updated access_token and refresh_token in config file")

	return nil
}
