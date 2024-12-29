package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	pkgE "github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type GetUserAnimeListParams[T types.NativeUserAnimeList | types.NativeAnimeList] struct {
	AnimeList           *T
	ListType            pkgE.UserAnimeListStatus
	Limit, Offset, Sort int
	Fields              []es.AnimeDetailField
}

func GetUserAnimeList[T types.NativeUserAnimeList | types.NativeAnimeList](p GetUserAnimeListParams[T]) error {
	fieldsStr := u.ConvertToCommaSeperatedString(p.Fields)

	sortOption := string(pkgE.UserAnimeListSortOptions()[p.Sort])
	parsedSortOptions, invalidFound := pkgE.ParseUserAnimeListSortOptions([]string{
		sortOption,
	})
	if invalidFound {
		return fmt.Errorf("Invalid Sort Option {%v}", sortOption)
	}

	// - ROUTE: /api/user/anime-list
	url := fmt.Sprintf("%s/user/anime-list?status=%s&limit=%v&offset=%v&sort=%s&fields=%v",
		e.ApiUrl,
        p.ListType,
		p.Limit,
		p.Offset,
		parsedSortOptions[0],
		fieldsStr,
	)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR GETTING USER ANIME LIST FROM SERVER****", err)
	}

	if err := json.NewDecoder(res.Body).Decode(p.AnimeList); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

	return nil
}
