package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
)


type GetUserAnimeFormDataParams struct {
    AnimeId int
	AnimeStatus *types.NativeUserAnimeStatus
}

func GetUserAnimeFormData(p GetUserAnimeFormDataParams) error {
	// - ROUTE: /api/user/anime/{animeID}/my_list_status
	url := fmt.Sprintf("%s/user/anime/%d/my_list_status",
		enums.ApiUrl,
        p.AnimeId,
	)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR GETTING USER ANIME DETAILS FROM SERVER****", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 && res.StatusCode < 600 {
		var v any
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			return fmt.Errorf("Json parsing error of anime-list \n %v", err)
		}
        jsonData, _ := json.MarshalIndent(v, "\t", " ")

		return fmt.Errorf("%s", jsonData)
	}

	if err := json.NewDecoder(res.Body).Decode(p.AnimeStatus); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

	return nil

}
