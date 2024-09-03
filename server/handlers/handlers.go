package handlers

import (
	"learning/server/utils"
	"log"
	"net/http"
)

// GET anime list
func GETAnimeList(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    log.Println(q)

    // get the data from MAL API
    data := utils.FetchAnimeList(utils.FetchAnimeListParams{
        Query: q.Get("q"),
        Limit: 10,
    });

    log.Print(data)
}

