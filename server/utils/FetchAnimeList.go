package utils

import (
	"encoding/json"
	"fmt"
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

type RES struct {
	Data []struct {
		Node struct {
			Id          int    `json:"id"`
			Title       string `json:"title"`
			MainPicture struct {
				Large  string `json:"large"`
				Medium string `json:"medium"`
			} `json:"main_picture"`
		} `json:"node"`
	} `json:"data"`
	Paging struct {
		Next string `json:"next"`
	} `json:"paging"`
}

func FetchAnimeList(p FetchAnimeListParams) RES {
	if p.Query == "" {
		return RES{}
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

	ret := RES{}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		log.Fatalf("ERROR in FetchAnimeList decoding json body \n %v", err)
	}

	return ret
}
