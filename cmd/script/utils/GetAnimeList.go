package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type GetAnimeListParams[T types.NativeAnimeList] struct {
	AnimeList     *T
	Query         string
	Limit, Offset int
	Fields        []es.AnimeDetailField
}

func GetAnimeList[T types.NativeAnimeList](p GetAnimeListParams[T]) error {
	encodedQuery := url.QueryEscape(p.Query)
	fieldsStr := u.ConvertToCommaSeperatedString(p.Fields)

	url := fmt.Sprintf("%s/anime-list?q=%s&limit=%v&offset=%v&fields=%v",
		enums.ApiUrl,
		encodedQuery,
		p.Limit,
		p.Offset,
		fieldsStr,
	)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR GETTING ANIME LIST FROM SERVER****", err)
	}

	if err := json.NewDecoder(res.Body).Decode(p.AnimeList); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

	return nil
}
