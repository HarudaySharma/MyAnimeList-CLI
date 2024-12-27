package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
)

type NativeUserDetailsParams struct {
    carry *types.NativeUserDetails
}
func GetUserDetails(p NativeUserDetailsParams) (error) {
	// - ROUTE: /api/anime/season/{year}/{season}?limit?offset?sort?fields
	url := fmt.Sprintf("%s/user/",
		enums.ApiUrl,
	)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR GETTING SEASONAL ANIME LIST FROM SERVER****", err)
	}

	if err := json.NewDecoder(res.Body).Decode(p.carry); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

    return nil

}
