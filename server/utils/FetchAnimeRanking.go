package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/server/config"
	"github.com/HarudaySharma/MyAnimeList-CLI/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/server/types"
)

type FetchAnimeRankingParams struct {
	Ranking       enums.Ranking
	Limit, Offset int
	Fields        string
}

// TODO: make types customizable for fields query
func FetchAnimeRanking(p FetchAnimeRankingParams) *types.NativeAnimeRanking {
	// handle params being non existent

	if p.Ranking == "" {
		return &types.NativeAnimeRanking{}
	}
	if p.Limit == 0 {
		p.Limit = enums.DEFAULT_LIMIT
	}
	if p.Limit > enums.MAX_LIMIT {
		p.Limit = enums.MAX_LIMIT
	}
	if p.Offset > enums.MAX_OFFSET {
		p.Offset = enums.DEFAULT_OFFSET
	}

	client := http.Client{}

	url := fmt.Sprintf("%s/anime/ranking?ranking_type=%s&limit=%v&offset=%v&fields=%s",
		config.C.MAL_API_URL,
		p.Ranking,
		p.Limit,
		p.Offset,
		p.Fields,
	)
	req := CreateHttpRequest("GET", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchAnimeList fetching from MAL API \n %v", err)
	}

	var ret types.MALAnimeRanking
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		log.Fatalf("ERROR in FetchAnimeList decoding json body \n %v", err)
	}

	defer res.Body.Close()

	return convertToNativeAnimeRankingType(&ret)
}

func convertToNativeAnimeRankingType(data *types.MALAnimeRanking) *types.NativeAnimeRanking {
	var parsedData types.NativeAnimeRanking

	for _, d := range data.Data {
		parsedData.Data = append(parsedData.Data, types.AnimeRankingDataNode{
			Node: types.AnimeListDataNode{
				ID:    d.Node.ID,
				Title: d.Node.Title,
			},
			Ranking: d.Ranking,
		})
	}
	parsedData.Paging = data.Paging

	return &parsedData
}
