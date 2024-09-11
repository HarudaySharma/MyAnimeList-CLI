package utils

import (
	"encoding/json"
	"fmt"
	"github.com/HarudaySharma/MyAnimeList-CLI/server/config"
	"github.com/HarudaySharma/MyAnimeList-CLI/server/types"
	"log"
	"net/http"
)

type FetchAnimeListParams struct {
	Query  string
	Limit  int8
	Offset int8
	Fields string
}

func FetchAnimeList(p FetchAnimeListParams) *types.NativeAnimeList {
	const DEFAULT_LIMIT = 10
	const DEFAULT_OFFSET = 0
	const MAX_LIMIT = 100
	const MAX_OFFSET = 99

	if p.Query == "" {
		return &types.NativeAnimeList{}
	}
	if p.Limit == 0 {
		p.Limit = DEFAULT_LIMIT
	}
	if p.Limit > MAX_LIMIT {
		p.Limit = MAX_LIMIT
	}
	if p.Offset > MAX_OFFSET {
		p.Offset = DEFAULT_OFFSET
	}

	// create a client
	client := http.Client{}

	url := fmt.Sprintf("%s/anime?q=%s&limit=%v&offset=%v",
		config.C.MAL_API_URL,
		p.Query,
		p.Limit,
		p.Offset,
	)
	req := CreateHttpRequest("GET", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchAnimeList fetching from MAL API \n %v", err)
	}

	ret := types.MALAnimeList{}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		log.Fatalf("ERROR in FetchAnimeList decoding json body \n %v", err)
	}

	defer res.Body.Close()

	return convertToNativeAnimeListType(&ret)
}

func convertToNativeAnimeListType(data *types.MALAnimeList) *types.NativeAnimeList {
	convertedData := types.NativeAnimeList{}

	for _, v := range data.Data {
		node := types.AnimeListDataNode{
			ID:    v.Node.ID,
			Title: v.Node.Title,
		}
		convertedData.Data = append(convertedData.Data, node)
	}
	convertedData.Paging = data.Paging

	return &convertedData
}
