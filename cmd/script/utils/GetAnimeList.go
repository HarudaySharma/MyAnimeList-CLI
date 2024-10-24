package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

func GetAnimeList[T any](animeList *T, query string, limit, offset int, fields []es.AnimeDetailField) error {
    encodedQuery := url.QueryEscape(query)
	fieldsStr := u.ConvertToCommaSeperatedString(fields)

	url := fmt.Sprintf("%s/anime-list?q=%s&limit=%v&offset=%v&fields=%v",
		enums.API_URL,
		encodedQuery,
		limit,
		offset,
		fieldsStr,
	)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR GETTING ANIME LIST FROM SERVER****", err)
	}

	if err := json.NewDecoder(res.Body).Decode(animeList); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

	return nil
}
