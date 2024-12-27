package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	srv "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/services"
	pkgE "github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
	pkgUtl "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

/*
NOTE: OAuth
  - user should first of all give permission to the app to modify/access it's mal data.
    1. send user to give permission page.
    2. get an authorization code obtained after permission grant
  - via callback url to (localhost) or,
  - let the user give it themself.
    3. save the auth_code and code_challenge/code_verifier.
    4. now whenever accessing the user data, use these auth code and
*/
func AuthCallback(w http.ResponseWriter, r *http.Request) {
	authCode := r.URL.Query().Get("code")
	if authCode == "" {
		log.Println("Invalid authCode")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// generate refresh and access tokens from this code.
	if err := pkgUtl.WriteConfigFile("auth_code", authCode); err != nil {
		log.Println("error saving auth_code in config file", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	config.C.MalAuthCode = authCode

	accessToken, refreshToken, err := srv.FetchOAuthTokens(srv.FetchOAuthTokensParams{
		GrantType:              srv.GrantTypeAuthorizationCode,
		ClientId:               config.C.MalClientId,
		CodeVerifier:           config.C.MalCodeVerifier,
		AuthCodeOrRefreshToken: authCode,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{
                "error": "%s",
                "hint": "retry giving access"
            }`,
			err)
		log.Println(err)
		return
	}

	log.Println("access_token: ", accessToken)
	log.Println("refresh_token: ", accessToken)

	if err := pkgUtl.WriteConfigFile("access_token", accessToken); err != nil {
		log.Println("error writing access_token to config file", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := pkgUtl.WriteConfigFile("refresh_token", refreshToken); err != nil {
		log.Println("error writing refresh_token to config file", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("saved access_token and refresh_token to config file")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{
            "message": "operation successfull, use the client application as desired"
        }`)

	return

}

func GETUserDetails(w http.ResponseWriter, r *http.Request) {
	if config.C.MalAuthCode == "" {
		config.C.MalAuthCode = pkgUtl.ReadConfigFile("auth_code")

		if config.C.MalAuthCode == "" {
			// user has not given permissions to access there data.
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{
                    "error": "authorization code not found",
                    "message": "please give authorization to access your mal data ",
                    "hint": "run mal-cli me login"
                }`)

			return
		}
	}

	data := srv.FetchUserDetails()
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{
                "message": "Internal server error"
            }`)

		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(jsonData))

	return
}

// GET https://api.myanimelist.net/v2/users/{user_name | @me}/animelist
func GETUserAnimeList(w http.ResponseWriter, r *http.Request) {
	// query params

	q := r.URL.Query()

	// anime list status
	userAnimeListStatus := q.Get("status")
	var parsedUALStatus pkgE.UserAnimeListStatus

	if userAnimeListStatus == "" {
		// return all the user anime list
		parsedUALStatus = ""
	}
	if userAnimeListStatus != "" {
		// parse the anime list type
		var valid bool
		parsedUALStatus, valid = pkgE.ParseUserAnimeListStatus(userAnimeListStatus)
		if !valid {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{
                "error": "invalid query param 'status'",
                "hint": "available sort options { %v }"
            }`,
				pkgUtl.ConvertToCommaSeperatedString(pkgE.UserAnimeListStatuses()))
			return
		}
	}

	// sort option
	sortOptions := strings.ReplaceAll(q.Get("sort"), " ", "")
	sortOptionArr := strings.Split(sortOptions, ",")

	parsedUALSortOptions, invalidFound := pkgE.ParseUserAnimeListSortOptions(sortOptionArr)
	if invalidFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{
                "error": "invalid query params 'sort'",
                "hint": "available sort options { %v }
            }`,
			pkgUtl.ConvertToCommaSeperatedString(pkgE.UserAnimeListSortOptions()))
		return
	}

	// parsing fields
	fields := strings.ReplaceAll(r.URL.Query().Get("fields"), " ", "")
	fieldArr := strings.Split(fields, ",")

	parsedFields, invalidFound := es.ParseDetailsField(fieldArr)
	if len(parsedFields) == 0 && invalidFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"invalid custom fields {available: %v }\"}", utils.ConvertToCommaSeperatedString(es.EveryDetailField()))
		return
	}

    // parsing limits and offsets
	limitStr := q.Get("limit")
	offsetStr := q.Get("offset")
	limit := 0
	offset := 0
	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(q.Get("limit")) // returns 0 if err
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrSyntax {
				fmt.Fprint(w, `{
                        "error": "invalid query params (invalid "limit"(0,1000) )"
                    }`)
				return
			}

			fmt.Fprint(w, `{
                    "error": "unexpected error"
                }`)
			return
		}
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(q.Get("offset"))
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrSyntax {
				fmt.Fprint(w, `{
                        "error": "invalid query params (invalid "offset"[0,1000) )}"
                    }`)
				return
			}

			fmt.Fprint(w, `{
                    "error": "unexpected error"
                }`)
			return
		}
	}

	data := srv.FetchUserAnimeList(srv.FetchUserAnimeListParams{
        Status: parsedUALStatus,
        Sort: parsedUALSortOptions,
        Fields: parsedFields,
        Limit: int16(limit),
        Offset: int16(offset),
    })

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{
                "message": "Internal server error"
            }`)

		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(jsonData))

}

// PUT https://api.myanimelist.net/v2/anime/{anime_id}/my_list_statusc
// DELETE https://api.myanimelist.net/v2/anime/{anime_id}/my_list_status
