package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	c "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type FetchSeasonalAnimeParams struct {
	Season e.AnimeSeason
	Year   string
	Sort   []e.SortOption
	Limit  int
	Offset int
	Fields []e.AnimeDetailField
}

func FetchSeasonalAnime(p FetchSeasonalAnimeParams) *t.NativeSeasonalAnime {
	if p.Limit == 0 {
		p.Limit = e.DEFAULT_SEARCH_LIST_SIZE
	}
	/* if p.Limit > e.MAX_LIMIT {
		p.Limit = e.MAX_LIMIT
	} */
	// doesn't need this actually
	/* if p.Offset > e.MAX_OFFSET {
		p.Offset = e.DEFAULT_OFFSET
	} */

	// create a client
	client := http.Client{}

	url := fmt.Sprintf("%s/anime/season/%s/%s?sort=%s&limit=%v&offset=%v&fields=%s",
		c.C.MalApiUrl,
		p.Year,
		p.Season,
		u.ConvertToCommaSeperatedString(p.Sort),
		p.Limit,
		p.Offset,
		u.ConvertToCommaSeperatedString(p.Fields),
	)

	log.Println(url)
	req := u.CreatePublicHttpRequest("GET", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchSeasonalAnimefetching from MAL API \n %v", err)
	}

	var ret t.MALSeasonalAnime
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		log.Fatalf("ERROR in FetchSeasonalAnimefetching decoding json body \n %v", err)
	}

	defer res.Body.Close()

	return convertToNativeSeasonalAnime(&ret)
}

func convertToNativeSeasonalAnime(data *t.MALSeasonalAnime) *t.NativeSeasonalAnime {
	convertedData := t.NativeSeasonalAnime{}

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

    b, _ := json.MarshalIndent(data, "\t", " ")
    fmt.Println(string(b))

	return &convertedData
}
