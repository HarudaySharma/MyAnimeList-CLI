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

type GetSeasonalAnimeParams[T types.NativeSeasonalAnime | types.NativeAnimeList] struct {
	AnimeList           *T
	Year                int
	Season              string
	Limit, Offset, Sort int
	Fields              []es.AnimeDetailField
}

func GetSeasonalAnime[T types.NativeSeasonalAnime | types.NativeAnimeList](p GetSeasonalAnimeParams[T]) error {
	fieldsStr := u.ConvertToCommaSeperatedString(p.Fields)

	sortOption := string(es.SortOptions()[p.Sort])
	parsedSortOptions, invalidFound := es.ParseSortOptions([]string{
		sortOption,
	})

	if invalidFound {
		return fmt.Errorf("Invalid Sort Option {%v}", sortOption)
	}

	// - ROUTE: /api/anime/season/{year}/{season}?limit?offset?sort?fields
	url := fmt.Sprintf("%s/anime/seasonal/%d/%s?limit=%v&offset=%v&sort=%s&fields=%v",
		enums.API_URL,
		p.Year,
		p.Season,
		p.Limit,
		p.Offset,
		parsedSortOptions[0],
		fieldsStr,
	)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR GETTING SEASONAL ANIME LIST FROM SERVER****", err)
	}

	if err := json.NewDecoder(res.Body).Decode(p.AnimeList); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

	return nil
}
