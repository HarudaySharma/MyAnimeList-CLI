package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type FetchUserAnimeStatusParams struct {
	AnimeID string
}

func FetchUserAnimeStatus(p FetchUserAnimeStatusParams) *t.NativeUserAnimeStatus {
	// create a client
	client := http.Client{}

    // PATCH https://api.myanimelist.net/v2/anime/{anime_id}/my_list_status
	url := fmt.Sprintf(`%s/anime/%s/my_list_status`,
		config.C.MalApiUrl,
		p.AnimeID,
	)
	log.Println(url)

	req := u.CreateUserHttpRequest(http.MethodPatch, url, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchUserAnimeStatus fetching from MAL API \n %v", err)
		log.Fatal(err)
		return &t.NativeUserAnimeStatus{}
	}

	ret := t.MALUserAnimeStatus{}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		log.Fatalf("ERROR in FetchUserAnimeStatus decoding json body \n %v", err)
	}

	defer res.Body.Close()

	return convertToNativeUserAnimeStatusType(&ret)
}

func convertToNativeUserAnimeStatusType(data *t.MALUserAnimeStatus) *t.NativeUserAnimeStatus {
	convertedData := t.NativeUserAnimeStatus{
		Status:             data.Status,
		Score:              data.Score,
		NumWatchedEpisodes: data.NumWatchedEpisodes,
		IsRewatching:       data.IsRewatching,
		UpdatedAt:          data.UpdatedAt,
	}

	return &convertedData
}
