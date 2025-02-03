package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/enums"
	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

type FetchUserAnimeListParams struct {
	Status enums.UserAnimeListStatus
	Sort   []enums.UserAnimeListSortOption
	Fields []es.AnimeDetailField
	Limit  int
	Offset int
}

func FetchUserAnimeList(p FetchUserAnimeListParams) *t.NativeUserAnimeList {
	if p.Limit == 0 {
		p.Limit = es.User_Default_List_Size
	}
	if p.Limit > es.User_Max_List_Size {
		p.Limit = es.User_Max_List_Size
	}
	if p.Offset < 0 || p.Offset > es.User_List_Max_Offset {
		p.Offset = es.User_List_Default_Offset
	}

	p.Fields = append(p.Fields, "list_status")
	fieldsStr := u.ConvertToCommaSeperatedString(p.Fields)

	sortStr := u.ConvertToCommaSeperatedString(p.Sort)

	// NOTE: if no status is provided then only mal api responses with all the user's anime list
	if p.Status == enums.ULS_ALL {
		p.Status = ""
	}

	// create a client
	client := http.Client{}

	url := fmt.Sprintf(`%s/users/@me/animelist?status=%s&sort=%s&limit=%d&offset=%d&fields=%s`,
		config.C.MalApiUrl,
		p.Status,
		sortStr,
		p.Limit,
		p.Offset,
		fieldsStr,
	)
	log.Println(url)

	req := u.CreateUserHttpRequest("GET", url, nil)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchUserAnimeList fetching from MAL API \n %v", err)
		log.Fatal(err)
		return &t.NativeUserAnimeList{}
	}

	ret := t.MALUserAnimeList{}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		// getting error if the query param (q = "word1 word2") i.e space b/w words is there
		log.Fatalf("ERROR in FetchUserAnimeList decoding json body \n %v", err)
	}

	defer res.Body.Close()

	return convertToNativeUserAnimeListType(&ret)
}

func convertToNativeUserAnimeListType(data *t.MALUserAnimeList) *t.NativeUserAnimeList {
	convertedData := t.NativeUserAnimeList{}

	for _, v := range data.Data {
		animeStatus := t.NativeUserAnimeStatus{
			Status:             v.AnimeStatus.Status,
			Score:              v.AnimeStatus.Score,
			NumWatchedEpisodes: v.AnimeStatus.NumEpisodesWatched,
			IsRewatching:       v.AnimeStatus.IsRewatching,
			UpdatedAt:          v.AnimeStatus.UpdatedAt,
		}

		node := t.UserAnimeListDataNode{
			AnimeStatus: animeStatus,
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
