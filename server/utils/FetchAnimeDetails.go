package utils

import (
	"encoding/json"
	"fmt"
	"learning/server/config"
	"log"
	"net/http"
)

// TODO: 
// - create a Enum such that user is in control of all the fields they want
// - Parse the field value before fetching
// - give users some default fields options or let them send custom fields
// - Default Options: "BasicDetails", "AdvancedDetails", "EveryDetail", "Custom"

func FetchAnimeDetails(animeId string, fields []AnimeDetailField) {
	client := http.Client{}

	url := fmt.Sprintf("%s/anime/%s?fields=%s",
		config.C.MAL_API_URL,
		animeId,
		fields,
	)

	req := CreateHttpRequest("GET", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var v any
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		log.Fatalf("ERROR in FetchAnimeDetails decoding json body \n %v", err)
	}

	body, _ := json.MarshalIndent(v, "", "\t")
	log.Print(string(body))

	return
}
