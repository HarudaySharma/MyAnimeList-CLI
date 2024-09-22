package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"

	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
)

func GetAnimeDetails[T any](animeDetails *T, animeId int, detailType string, fields []es.AnimeDetailField) error {
	fieldsStr := utils.ConvertToCommaSeperatedString(fields)

	url := fmt.Sprintf("%s/anime/%d?detail_type=%s&fields=%s",
		enums.API_URL,
		animeId,
		detailType,
		fieldsStr,
	)

	res, err := http.Get(url)
	if err != nil {
		panic("error getting anime list")
	}

	if err := json.NewDecoder(res.Body).Decode(animeDetails); err != nil {
		return fmt.Errorf("Json parsing error of animeDetails \n %v", err)
	}

	return nil
}
