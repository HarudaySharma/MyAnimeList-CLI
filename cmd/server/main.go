package main

import (
	"github.com/HarudaySharma/MyAnimeList-CLI/server"
	"log"
)

func main() {
    log.SetFlags(log.Ltime)

    server.StartServer("3000")
}
