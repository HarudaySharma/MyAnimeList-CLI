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

type FetchAnimeListParams struct {
	Query  string
	Limit  int8
	Offset int8
	Fields string
}

func FetchAnimeList(p FetchAnimeListParams) *types.NativeAnimeList {

	if p.Query == "" {
		return &types.NativeAnimeList{}
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
