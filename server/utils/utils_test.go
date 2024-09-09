package utils

import (
	"encoding/json"
	"learning/server/config"
	"testing"
)

func Test_FetchAnimeList(t *testing.T) {
	config.LoadConfig()

	ret := FetchAnimeList(FetchAnimeListParams{
		Query: "Vinland",
		Limit: 5,
	})

    data, err := json.MarshalIndent(ret, "", "\t")
    if err != nil {
        t.Fatal(err)
    }
    t.Log(string(data))

}
