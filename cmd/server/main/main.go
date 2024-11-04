package main

import (
	"log"
	"os"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server"
)

func main() {
	log.SetFlags(log.Ltime)

    if len(os.Args) < 2 {
        log.Panic("SPECIFY A PORT")
        os.Exit(1)
    }

    port := os.Args[1]
	server.StartServer(port)
}
