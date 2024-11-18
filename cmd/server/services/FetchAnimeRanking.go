package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type FetchAnimeRankingParams struct {
	Ranking       e.Ranking
	Limit, Offset int
	Fields        []e.AnimeDetailField
}

func FetchAnimeRanking(p FetchAnimeRankingParams) *t.NativeAnimeRanking {
	// handle params being non existent

	log.Println("fetching anime by ranking")
	if p.Ranking == "" {
		return &t.NativeAnimeRanking{}
	}
	if p.Limit == 0 {
		p.Limit = e.DEFAULT_SEARCH_LIST_SIZE
	}
	if p.Limit > e.MAX_RANKING_SEARCH_LIST_SIZE {
        // NOTE: Seems we can request at most 500 data nodes at a time from mal api at this endpoint
		return &t.NativeAnimeRanking{}
	}
	// doesn't need this actually
	/* if p.Offset > e.MAX_OFFSET {
		p.Offset = e.DEFAULT_OFFSET
	} */

	client := http.Client{}

	fields := u.ConvertToCommaSeperatedString(p.Fields)
	url := fmt.Sprintf("%s/anime/ranking?ranking_type=%s&limit=%v&offset=%v&fields=%s",
		config.C.MalApiUrl,
		p.Ranking,
		p.Limit,
		p.Offset,
		fields,
	)

	log.Println(url)
	req := u.CreateHttpRequest("GET", url)
	//log.Println("Fetching from URL:", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchAnimeList fetching from MAL API \n %v", err)
	}

	var ret t.MALAnimeRanking
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		log.Fatalf("ERROR in FetchAnimeList decoding json body \n %v", err)
	}

	defer res.Body.Close()

	//log.Printf("CustomFields for ID %d: %+v\n", ret.Data[0].Node.ID, ret.Data[0].Node.CustomFields) // CustomFields correctly populated

	return convertToNativeAnimeRankingType(&ret)
}

func convertToNativeAnimeRankingType(data *t.MALAnimeRanking) *t.NativeAnimeRanking {
	var parsedData t.NativeAnimeRanking
	for _, d := range data.Data {
		node := t.AnimeListDataNode{
			ID:           d.Node.ID,
			Title:        d.Node.Title,
            CustomFields: d.Node.CustomFields, // Still keeping the raw custom fields
		}
        node.CustomFields["main_picture"] = d.Node.MainPicture

		parsedData.Data = append(parsedData.Data, t.AnimeRankingDataNode{
            Node: node,
			Ranking: d.Ranking,
		})
	}
	parsedData.Paging = data.Paging

	return &parsedData
}
