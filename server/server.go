package server

import (
	"learning/server/handlers"
	"log"
	"net/http"
)

func StartServer(port string) {
	log.Printf("server running on http://localhost:%s", port)

	// create request routesz
	http.HandleFunc("/api/anime-list", handlers.GETAnimeList)

    // NOTE: don't remove the "/" at the end of endpoint (for dynamic routing)
	http.HandleFunc("/api/anime/", handlers.GETAnimeDetails)

	// SERVER
	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	log.Fatal(s.ListenAndServe())
}
