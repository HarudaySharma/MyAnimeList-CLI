package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
)

type DeleteUserAnimeStatusParams struct {
	AnimeId int
}

func DeleteUserAnimeStatus(p DeleteUserAnimeStatusParams) error {
	// - ROUTE: /api/user/anime/{animeID}/my_list_status
	url := fmt.Sprintf("%s/user/anime/%d/my_list_status",
		enums.ApiUrl,
		p.AnimeId,
	)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR CREATING DELETE REQUEST****", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR DELETING USER ANIME STATUS FROM SERVER****", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 && res.StatusCode < 600 {
		var v any
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			return fmt.Errorf("Json parsing error \n %v", err)
		}
		jsonData, _ := json.MarshalIndent(v, "\t", " ")

		return fmt.Errorf("%s", jsonData)
	}

	return nil

}
