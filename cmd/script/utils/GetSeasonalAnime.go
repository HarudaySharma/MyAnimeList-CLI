package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

//- ROUTE: /api/anime/season/{year}/{season}?limit?offset?sort?fields
func GetSeasonalAnime[T any](animeList *T,  year int, season string, limit, offset int, fields []es.AnimeDetailField) error {
	fieldsStr := u.ConvertToCommaSeperatedString(fields)
    var seasonalAnime *types.NativeSeasonalAnime

	url := fmt.Sprintf("%s/anime/seasonal/%d/%s?limit=%v&offset=%v&fields=%v",
		enums.API_URL,
		year,
		season,
		limit,
		offset,
		fieldsStr,
	)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR GETTING SEASONAL ANIME LIST FROM SERVER****", err)
	}

	if err := json.NewDecoder(res.Body).Decode(seasonalAnime); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

    bt, _ := json.MarshalIndent(seasonalAnime, "", "\t")
    fmt.Println(string(bt))

	return nil
}
