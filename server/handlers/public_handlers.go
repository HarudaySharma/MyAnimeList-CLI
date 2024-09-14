package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/server/utils"
)

/*
GET anime list
  - ROUTE: /api/anime/anime-list?q="query.."&limit="[1, 100]"&offset="[0, 99]"
*/
func GETAnimeList(w http.ResponseWriter, r *http.Request) {
	log.Println("*****GETAnimeList Handler called*****")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "{\"error\": \"Only GET request is allowed\"}")
		return
	}

	q := r.URL.Query()

	if q.Get("q") == "" {
		fmt.Fprint(w, "{\"error\": \"invalid query params (\"q\" not provided)\"}")
		return
	}

	limitStr := q.Get("limit")
	offsetStr := q.Get("offset")
	limit := 0
	offset := 0
	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(q.Get("limit")) // returns 0 if err
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrSyntax {
				fmt.Fprint(w, "{\"error\": \"invalid query params (invalid \"limit\"(0,100)}")
				return
			}

			fmt.Fprint(w, "{\"error\": \"unexpected error\"}")
			return
		}
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(q.Get("offset"))
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrSyntax {
				fmt.Fprint(w, "{\"error\": \"invalid query params (invalid \"offset\"[0,100)\"}")
				return
			}

			fmt.Fprint(w, "{\"error\": \"unexpected error\"}")
			return
		}
	}

	data := utils.FetchAnimeList(utils.FetchAnimeListParams{
		Query:  q.Get("q"),
		Limit:  int8(limit),
		Offset: int8(offset),
	})

	//utils.PrintJSON(data)
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"message\": \"Internal server error\"}")
		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(jsonData))

	return
}

/*
GET anime details
  - ROUTE: /api/anime/{animeId}?detail_type="Basic"|"Advanced"|"Custom"&fields="....."
  - if detail_type is "custom" then the custom query param should be filled too
  - Returns blank json object with id: 0 for invalid animeId
*/
func GETAnimeDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("*****GETAnimeDetails Handler called*****")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "{\"error\": \"Only GET request is allowed\"}")
		return
	}

	// ****animeId PARSING****

	p := r.URL.Path
	pathSegments := strings.Split(p, "/")

	if len(pathSegments) != 4 {
		if len(pathSegments) != 5 || pathSegments[len(pathSegments)-1] != "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	var animeId string
	if pathSegments[len(pathSegments)-1] == "" {
		// FOR: /api/anime/{animeID}/
		animeId = pathSegments[len(pathSegments)-2]
	} else {
		// FOR: /api/anime/{animeID}
		animeId = pathSegments[len(pathSegments)-1]
	}
	log.Printf("animeId: %s", animeId)

	// ****QUERY PARAMS****

	var data interface{}
	detailType := r.URL.Query().Get("detail_type")
	switch strings.ToLower(detailType) {

	case "":
		data = utils.FetchAnimeDetails(animeId, enums.EveryDetailField())

	case "basic":
		data = utils.FetchAnimeDetails(animeId, enums.BasicDetailFields())

	case "advanced":
		data = utils.FetchAnimeDetails(animeId, enums.AdvancedDetailFields())

	case "custom":
		// get the "custom" query param
		fields := strings.ReplaceAll(r.URL.Query().Get("fields"), " ", "")
		fieldArr := strings.Split(fields, ",")

		parsedFields, invalidFound := enums.ParseDetailsField(fieldArr)
		if len(parsedFields) == 0 && invalidFound {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "{\"error\": \"invalid custom fields\"}")
			return
		}
		data = utils.FetchAnimeDetails(animeId, parsedFields)

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{\"error\": \"invalid detail_type\"}")
		return
	}

	jsonData, err := json.Marshal(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": \"JSON parsing failed\"}")
		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(jsonData))

	return
}

/*
GET anime ranking
  - ROUTE: /api/anime/ranking?ranking_type&limit&offset&fields
*/
func GETAnimeRanking(w http.ResponseWriter, r *http.Request) {
	log.Println("*****GETAnimeRanking Handler called*****")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "{\"error\": \"Only GET request is allowed\"}")
		return
	}


	q := r.URL.Query()

	rankingType := q.Get("ranking_type")
	rankingTypeParsed, ok := enums.ParseAnimeRaking(rankingType)
	if rankingType == "" || !ok {
		fmt.Fprint(w, "{\"error\": \"invalid query params \"ranking_type\"\"}")
		return
	}

	limitStr := q.Get("limit")
	offsetStr := q.Get("offset")
	limit := 0
	offset := 0
	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(q.Get("limit")) // returns 0 if err
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrSyntax {
				fmt.Fprint(w, "{\"error\": \"invalid query params (invalid \"limit\"(0,100)}")
				return
			}

			fmt.Fprint(w, "{\"error\": \"unexpected error\"}")
			return
		}
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(q.Get("offset"))
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrSyntax {
				fmt.Fprint(w, "{\"error\": \"invalid query params (invalid \"offset\"[0,100)\"}")
				return
			}

			fmt.Fprint(w, "{\"error\": \"unexpected error\"}")
			return
		}
	}

	// parsing fields
	fields := strings.ReplaceAll(r.URL.Query().Get("fields"), " ", "")
	fieldArr := strings.Split(fields, ",")

	parsedFields, invalidFound := enums.ParseDetailsField(fieldArr)
	if len(parsedFields) == 0 && invalidFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{\"error\": \"invalid custom fields\"}")
		return
	}

	// fetching anime ranking
	data := utils.FetchAnimeRanking(utils.FetchAnimeRankingParams{
		Ranking: rankingTypeParsed,
		Limit:   limit,
		Offset:  offset,
		Fields:  parsedFields,
	})

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": \"JSON parsing failed\"}")
		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(jsonData))

	return
}

/*
GET Seasonal Anime
  - ROUTE: /api/anime/season/{year}/{season}?limit?offset?sort?fields
*/
func GETSeasonalAnime(w http.ResponseWriter, r *http.Request) {
	log.Println("*****GETSeasonalAnime Handler called*****")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "{\"error\": \"Only GET request is allowed\"}")
		return
	}

	// ****PATH Params****

	p := r.URL.Path
	pathSegments := strings.Split(p, "/")

	if len(pathSegments) != 6 {
		if len(pathSegments) != 7 || pathSegments[len(pathSegments)-1] != "" {
            log.Printf("here")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	var season, year string
	if pathSegments[len(pathSegments)-1] == "" {
		// FOR: /api/anime/{year}/{season}/
		season = pathSegments[len(pathSegments)-2]
		year = pathSegments[len(pathSegments)-3]
	} else {
		// FOR: /api/anime/{year}/{season}
		season = pathSegments[len(pathSegments)-1]
		year = pathSegments[len(pathSegments)-2]
	}

	// ****QUERY PARAMS****

	q := r.URL.Query()
	limitStr := q.Get("limit")
	offsetStr := q.Get("offset")
	limit := 0
	offset := 0
	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(q.Get("limit")) // returns 0 if err
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrSyntax {
				fmt.Fprint(w, "{\"error\": \"invalid query params (invalid \"limit\"(0,100)}")
				return
			}

			fmt.Fprint(w, "{\"error\": \"unexpected error\"}")
			return
		}
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(q.Get("offset"))
		if err != nil {
			if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrSyntax {
				fmt.Fprint(w, "{\"error\": \"invalid query params (invalid \"offset\"[0,100)\"}")
				return
			}

			fmt.Fprint(w, "{\"error\": \"unexpected error\"}")
			return
		}
	}

	// parsing fields
	fields := strings.ReplaceAll(r.URL.Query().Get("fields"), " ", "")
	fieldArr := strings.Split(fields, ",")

	parsedFields, invalidFound := enums.ParseDetailsField(fieldArr)
	if invalidFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{\"error\": \"invalid custom fields\"}")
		return
	}

	// parsing sort
	// valid values: anime_score, anime_num_list_users
	// sort order: descending
	sort := q.Get("sort")

	if sort != "" && sort != "anime_score" && sort != "anime_num_list_users" {
		fmt.Fprint(w, "{\"error\": \"invalid query params \"sort{anime_num_list_users, anime_score}\"\"}")
		return
	}



}
