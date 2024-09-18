package server

import (
	"log"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/handlers"
)

func StartServer(port string) {
	log.Printf("server running on http://localhost:%s", port)

	// create request routez
	http.HandleFunc("/api/anime-list", handlers.GETAnimeList)
    // NOTE: don't remove the "/" at the end of endpoint (for dynamic routing)
    http.HandleFunc("/api/anime/", handlers.GETAnimeDetails)
    http.HandleFunc("/api/anime/ranking", handlers.GETAnimeRanking)
    // NOTE: don't remove the "/" at the end of endpoint (for dynamic routing)
    http.HandleFunc("/api/anime/seasonal/", handlers.GETSeasonalAnime)


	// SERVER
	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	log.Fatal(s.ListenAndServe())
}
