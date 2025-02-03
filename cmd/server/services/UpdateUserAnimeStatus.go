package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	urlPkg "net/url"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type UpdateUserAnimeStatusParams struct {
	AnimeID     string
	AnimeStatus *types.NativeUserAnimeStatus
}

func UpdateUserAnimeStatus(p UpdateUserAnimeStatusParams) *t.NativeUserAnimeStatus {
	// create a client
	client := http.Client{}

	// PUT https://api.myanimelist.net/v2/anime/{anime_id}/my_list_status
	url := fmt.Sprintf(`%s/anime/%s/my_list_status`,
		config.C.MalApiUrl,
		p.AnimeID,
	)
	log.Println(url)

    // Prepare form data
	formData := urlPkg.Values{}
	if p.AnimeStatus.Status != "" {
		formData.Set("status", string(p.AnimeStatus.Status))
	}
	if p.AnimeStatus.Score != -1 {
		formData.Set("score", fmt.Sprintf("%d", p.AnimeStatus.Score))
	}
    if p.AnimeStatus.NumWatchedEpisodes != -1 {
        formData.Set("num_watched_episodes", fmt.Sprintf("%d",p.AnimeStatus.NumWatchedEpisodes))
    }

    log.Println(formData)
	req := u.CreateUserHttpRequest(http.MethodPatch, url, bytes.NewReader([]byte(formData.Encode())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in Updating User Anime Status form MAL API \n %v", err)
		log.Fatal(err)
		return &t.NativeUserAnimeStatus{}
	}

	ret := t.MALUserAnimeStatus{}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		log.Fatalf("ERROR in UpdateUserAnimeStatus decoding json body \n %v", err)
	}

	defer res.Body.Close()

	return convertToNativeUserAnimeStatusType(&ret)
}
