package utils

import (
	"fmt"
	"io"
	"learning/server/config"
	"log"
	"net/http"
)

type FetchAnimeListParams struct {
	Query  string
	Limit  int8
	Offset int
	Fields string
}

func FetchAnimeList(p FetchAnimeListParams) []byte {
	if p.Query == "" {
		return nil
	}

	// create a client
	client := http.Client{}

	url := fmt.Sprintf("%s/q=%s&limit=%v&offset=%v&client_auth=%s",
		config.C.MAL_API_URL,
		p.Query,
		p.Limit,
		p.Offset,
        config.C.CLIENT_ID,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("ERROR in FetchAnimeList \n %v", err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchAnimeList fetching from MAL API \n %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ERROR in FetchAnimeList reading body \n %v", err)
	}

	return body
}
