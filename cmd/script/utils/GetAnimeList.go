package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

func GetAnimeList[T any](animeList *T, query string, limit, offset int, fields []es.AnimeDetailField) error {
	fieldsStr := u.ConvertToCommaSeperatedString(fields)

	url := fmt.Sprintf("%s/anime-list?q=%s&limit=%v&offset=%v&fields=%v",
		enums.API_URL,
		query,
		limit,
		offset,
		fieldsStr,
	)

	res, err := http.Get(url)
	if err != nil {
		panic("error getting anime list")
	}

	if err := json.NewDecoder(res.Body).Decode(animeList); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

	return nil
}
