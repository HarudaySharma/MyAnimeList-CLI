package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type FetchAnimeListParams struct {
	Query  string
	Limit  int8
	Offset int8
	Fields []e.AnimeDetailField
}

func FetchAnimeList(p FetchAnimeListParams) *t.NativeAnimeList {
	if p.Query == "" {
		return &t.NativeAnimeList{}
	}
	if p.Limit == 0 {
		p.Limit = e.DEFAULT_SEARCH_LIST_SIZE
	}
	if p.Limit > e.MAX_SEARCH_LIST_SIZE {
		p.Limit = e.MAX_SEARCH_LIST_SIZE
	}
	// doesn't need this actually
	/* if p.Offset > e.MAX_OFFSET {
		p.Offset = e.DEFAULT_OFFSET
	} */

	// create a client
	client := http.Client{}

	encodedQuery := url.QueryEscape(p.Query)
	fieldsStr := u.ConvertToCommaSeperatedString(p.Fields)
	url := fmt.Sprintf("%s/anime?q=%s&limit=%v&offset=%v&fields=%v",
		config.C.MalApiUrl,
		encodedQuery,
		p.Limit,
		p.Offset,
		fieldsStr,
	)
	req := u.CreatePublicHttpRequest("GET", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchAnimeList fetching from MAL API \n %v", err)
	}

	ret := t.MALAnimeList{}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		// getting error if the query param (q = "word1 word2") i.e space b/w words is there
		log.Fatalf("ERROR in FetchAnimeList decoding json body \n %v", err)
	}

	defer res.Body.Close()

	return convertToNativeAnimeListType(&ret)
}

func convertToNativeAnimeListType(data *t.MALAnimeList) *t.NativeAnimeList {
	convertedData := t.NativeAnimeList{}

	for _, v := range data.Data {
		node := t.AnimeListDataNode{
			ID:           v.Node.ID,
			Title:        v.Node.Title,
			CustomFields: v.Node.CustomFields,
		}
        node.CustomFields["main_picture"] = v.Node.MainPicture
		convertedData.Data = append(convertedData.Data, node)
	}
	convertedData.Paging = data.Paging

	return &convertedData
}
