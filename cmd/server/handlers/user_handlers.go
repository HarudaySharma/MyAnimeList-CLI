package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")
	if authCode == "" {
		log.Println("Invalid authCode")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// generate refresh and access tokens from this code.
	if err := utils.WriteConfigFile("auth_code", authCode); err != nil {
		log.Println("error saving auth_code in config file", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	config.C.MalAuthCode = authCode

	// get the auth tokens
	values := url.Values{}
	values.Add("client_id", config.C.MalClientId)
	values.Add("code_verifier", config.C.MalCodeVerifier)
	values.Add("code", config.C.MalAuthCode)
	values.Add("grant_type", "authorization_code")

	resp, err := http.PostForm("https://myanimelist.net/v1/oauth2/token", values)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var data struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("error decoding body into valid json", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
    defer resp.Body.Close()

	log.Println("access_token: ", data.AccessToken)
	log.Println("refresh_token: ", data.RefreshToken)

	if err := utils.WriteConfigFile("access_token", data.AccessToken); err != nil {
		log.Println("error writing access_token to config file", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := utils.WriteConfigFile("refresh_token", data.RefreshToken); err != nil {
		log.Println("error writing refresh_token to config file", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("saved access_token and refresh_token to config file")

	w.WriteHeader(http.StatusOK)
	return

}

func GETUserInfo(w http.ResponseWriter, r *http.Request) {
	if config.C.MalAuthCode == "" {
		config.C.MalAuthCode = utils.ReadConfigFile("auth_code")

		if config.C.MalAuthCode == "" {
			// user has not given permissions to access there data.
			fmt.Fprint(w, "please give permission first")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	client := http.Client{}
	req := utils.CreateUserHttpRequest("GET", "https://api.myanimelist.net/v2/users/@me")
	client.Do(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in GETUserInfo fetching from MAL API \n %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var v any
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		log.Fatalf("ERROR decoding json body \n %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
    defer resp.Body.Close()

	if checkInvalidAccessToken(v) {
		// TODO: refresh access token
		log.Printf("access_token invalid or expired probably %+v", v)
        return
	}

	jsonData, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"message\": \"Internal server error\"}")
		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(jsonData))

	return
}

func checkInvalidAccessToken(v any) bool {
	if errMap, ok := v.(map[string]interface{}); ok {
		if error, ok := errMap["error"].(string); ok && error == "invalid_token" {
			return true
		}
	}
	return false
}
