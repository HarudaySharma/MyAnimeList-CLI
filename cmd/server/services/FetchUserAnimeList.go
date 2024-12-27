package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	e "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

const (
	default_list_size int16 = 100
	max_list_size     int16 = 1000
	default_offset    int16 = 0
	max_offset        int16 = 1000
)

type FetchUserAnimeListParams struct {
	Status enums.UserAnimeListStatus
	Sort   []enums.UserAnimeListSortOption
	Fields []e.AnimeDetailField
	Limit  int16
	Offset int16
}

func FetchUserAnimeList(p FetchUserAnimeListParams) *t.NativeUserAnimeList {
	if p.Limit == 0 {
		p.Limit = default_list_size
	}
	if p.Limit > max_list_size {
		p.Limit = e.MAX_SEARCH_LIST_SIZE
	}
	// doesn't need this actually
	/* if p.Offset > e.MAX_OFFSET {
		p.Offset = e.DEFAULT_OFFSET
	} */

	p.Fields = append(p.Fields, "list_status")
	fieldsStr := u.ConvertToCommaSeperatedString(p.Fields)

	// create a client
	client := http.Client{}

	url := fmt.Sprintf(`%s/users/@me/animelist?status=%s&sort=%s&limit=%d&offset=%d&fields=%s`,
		config.C.MalApiUrl,
		p.Status,
		u.ConvertToCommaSeperatedString(p.Sort),
		p.Limit,
		p.Offset,
		fieldsStr,
	)
	req := u.CreateUserHttpRequest("GET", url)
	fmt.Println(url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchUserDetails fetching from MAL API \n %v", err)
		log.Fatal(err)
		return &t.NativeUserAnimeList{}
	}

	ret := t.MALUserAnimeList{}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		// getting error if the query param (q = "word1 word2") i.e space b/w words is there
		log.Fatalf("ERROR in FetchAnimeList decoding json body \n %v", err)
	}

	defer res.Body.Close()

	return convertToNativeUserAnimeListType(&ret)
}

func convertToNativeUserAnimeListType(data *t.MALUserAnimeList) *t.NativeUserAnimeList {
	convertedData := t.NativeUserAnimeList{}

	for _, v := range data.Data {
		node := t.UserAnimeListDataNode{
			ListStatus: v.ListStatus,
			Node: t.AnimeListDataNode{
				ID:           v.Node.ID,
				Title:        v.Node.Title,
				CustomFields: v.Node.CustomFields,
			},
		}
		/* if node.Node.CustomFields == nil {
			node.Node.CustomFields = make(map[string]interface{})
		} */
		node.Node.CustomFields["main_picture"] = v.Node.MainPicture
		convertedData.Data = append(convertedData.Data, node)
	}
	convertedData.Paging = data.Paging

	return &convertedData
}
