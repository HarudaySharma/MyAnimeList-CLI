package utils

import (
	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	"log"
	"net/http"
)

func CreateHttpRequest(method, url string) (*http.Request) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalf("ERROR creating new request \n %v", err)
		return nil
	}

	// don't use Authorization and X-MAL-CLIENT-ID in conjuction
	//req.Header.Add("Authorization", "Bearer " +  config.C.MAL_CLIENT_ID)
	req.Header.Add("X-MAL-CLIENT-ID", config.C.MalClientId)

	return req
}
