package server

import (
	"learning/server/config"
	"learning/server/handlers"
	"log"
	"net/http"
)


func StartServer(port string) {
	log.Printf("server running on http://localhost:%s", port)

	// create request routesz
    http.HandleFunc("/anime-list", handlers.GETAnimeList)

	// SERVER
	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

    config.LoadConfig()
	log.Fatal(s.ListenAndServe())
}
