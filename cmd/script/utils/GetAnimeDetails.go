package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"

	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
)

type GetAnimeDetailsParams[T types.AnimeDetails] struct {
	AnimeDetails *T
	AnimeId      int
	DetailType   string
	Fields       []es.AnimeDetailField
}

func GetAnimeDetails[T types.AnimeDetails](p GetAnimeDetailsParams[T]) error {
	fieldsStr := utils.ConvertToCommaSeperatedString(p.Fields)

	url := fmt.Sprintf("%s/anime/%d?detail_type=%s&fields=%s",
		enums.ApiUrl,
		p.AnimeId,
		p.DetailType,
		fieldsStr,
	)

	res, err := http.Get(url)
	if err != nil {
		panic("error getting anime list")
	}

	if err := json.NewDecoder(res.Body).Decode(p.AnimeDetails); err != nil {
		return fmt.Errorf("Json parsing error of animeDetails \n %v", err)
	}

	return nil
}
