package server

import (
	"log"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/handlers"
)

func StartServer(port string) {
	log.Printf("server running on http://localhost:%s", port)

	// create request routez

	http.HandleFunc("GET /api/anime-list", handlers.GETAnimeList)
    // NOTE: don't remove the "/" at the end of endpoint (for dynamic routing)
    http.HandleFunc("GET /api/anime/", handlers.GETAnimeDetails)
    http.HandleFunc("GET /api/anime/ranking", handlers.GETAnimeRanking)
    // NOTE: don't remove the "/" at the end of endpoint (for dynamic routing)
    http.HandleFunc("GET /api/anime/seasonal/", handlers.GETSeasonalAnime)


    // user specific routes
	http.HandleFunc("GET /api/callback", handlers.AuthCallback)
	http.HandleFunc("GET /api/auth", handlers.AuthCallback)
	http.HandleFunc("GET /api/user", handlers.GETUserDetails)
	http.HandleFunc("GET /api/user/anime-list", handlers.GETUserAnimeList)

    // Expected path: /api/user/anime/{animeid}/my_list_status
    http.HandleFunc("GET /api/user/anime/", handlers.GETUserAnimeStatus)
    http.HandleFunc("PATCH /api/user/anime/", handlers.PATCHUserAnimeStatus)
    http.HandleFunc("DELETE /api/user/anime/", handlers.DELETEUserAnimeStatus)

	// SERVER
	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	log.Fatal(s.ListenAndServe())
}
