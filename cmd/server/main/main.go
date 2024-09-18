package main

import (
	"log"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server"
)

func main() {
	log.SetFlags(log.Ltime)

	server.StartServer("3000")
}
