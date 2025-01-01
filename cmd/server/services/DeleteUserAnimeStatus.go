package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type DeleteUserAnimeStatusParams struct {
	AnimeID string
}

func DeleteUserAnimeStatus(p DeleteUserAnimeStatusParams) error {
	// create a client
	client := http.Client{}

	// DELETE https://api.myanimelist.net/v2/anime/{anime_id}/my_list_status
	url := fmt.Sprintf(`%s/anime/%s/my_list_status`,
		config.C.MalApiUrl,
		p.AnimeID,
	)
	log.Println(url)

	req := u.CreateUserHttpRequest(http.MethodDelete, url, nil)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in DeleteUserAnimeStatus fetching from MAL API \n %v", err)
		log.Fatal(err)
		return err
	}

	if res.StatusCode == http.StatusNotFound {
		return fmt.Errorf(`specified anime does not exist in user's anime list`)
	}

	return nil
}
