package utils

import (
	"encoding/json"
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

func FetchAnimeList(p FetchAnimeListParams) any {
	if p.Query == "" {
		return nil
	}

	// create a client
	client := http.Client{}

	url := fmt.Sprintf("%s/anime?q=%s&limit=%v&offset=%v",
		config.C.MAL_API_URL,
		p.Query,
		p.Limit,
		p.Offset,
	)
    req := CreateHttpRequest("GET", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchAnimeList fetching from MAL API \n %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ERROR in FetchAnimeList reading body \n %v", err)
	}


    log.Println(string(body))
    var ret any;
    if err := json.Unmarshal(body, ret); err != nil {
        log.Fatal(err)
    }
	return ret
}
