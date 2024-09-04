package handlers

import (
	"encoding/json"
	"fmt"
	"learning/server/utils"
	"net/http"
	"strconv"
)

// GET anime list
func GETAnimeList(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if q.Get("q") == "" {
		fmt.Fprint(w, "{\"error\": \"invalid query params (\"q\" not provided)}\"")
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
				fmt.Fprint(w, "{\"error\": \"invalid query params (invalid \"offset\"[0,100)}\"")
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
