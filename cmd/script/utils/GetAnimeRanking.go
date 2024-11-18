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

type GetAnimeRankingParams[T types.NativeAnimeRanking | types.NativeAnimeList] struct {
	AnimeList           *T
	RankingType         string
	Limit, Offset, Sort int
	Fields              []es.AnimeDetailField
}

func GetAnimeRanking[T types.NativeAnimeRanking | types.NativeAnimeList](p GetAnimeRankingParams[T]) error {
	fieldsStr := u.ConvertToCommaSeperatedString(p.Fields)

	_, invalid := es.ParseAnimeRaking(p.RankingType)
	if invalid {
		return fmt.Errorf("Invalid Ranking Type Option {%v}", p.RankingType)
	}
	// ROUTE: /api/anime/ranking?ranking_type&limit&offset&fields
	url := fmt.Sprintf("%s/anime/ranking?ranking_type=%s&limit=%d&offset=%d&fields=%s",
		enums.ApiUrl,
		p.RankingType,
		p.Limit,
		p.Offset,
		fieldsStr,
	)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR GETTING RANKING ANIME LIST FROM SERVER****", err)
	}

	if err := json.NewDecoder(res.Body).Decode(p.AnimeList); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

	return nil
}
