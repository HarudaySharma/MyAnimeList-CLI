package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/HarudaySharma/MyAnimeList-CLI/server/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	utils.PrintJSON(data)
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"message\": \"Internal server error\"}")
		return
	}

	w.Header().Add("content-type", "application/json")
	fmt.Fprint(w, string(jsonData))

	return
}

/*
GET anime details
  - ROUTE: /api/anime/{animeId}?detailType="Basic"|"Advanced"|"Custom"&custom="....."
  - if detailType is "custom" then the custom query param should be filled too
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
	detailType := r.URL.Query().Get("detailType")
	switch strings.ToLower(detailType) {

	case "":
		data = utils.FetchAnimeDetails(animeId, utils.EveryDetailField())

	case "basic":
		data = utils.FetchAnimeDetails(animeId, utils.BasicDetailFields())

	case "advanced":
		data = utils.FetchAnimeDetails(animeId, utils.AdvancedDetailFields())

	case "custom":
		// get the "custom" query param
		fields := strings.ReplaceAll(r.URL.Query().Get("custom"), " ", "")
		fieldArr := strings.Split(fields, ",")

		parsedFields, invalidFound := utils.ParseDetailsField(fieldArr)
		if len(parsedFields) == 0 && invalidFound {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "{\"error\": \"invalid custom fields\"}")
			return
		}
		data = utils.FetchAnimeDetails(animeId, parsedFields)

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{\"error\": \"invalid detailType\"}")
		return
	}

	jsonData, err := json.Marshal(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": \"JSON parsing failed\"}")
		return
	}

	fmt.Fprint(w, string(jsonData))
}
