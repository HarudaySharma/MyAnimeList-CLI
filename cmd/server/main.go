package main

import (
	"learning/server"
	"log"
)

func main() {
    log.SetFlags(log.Ltime)

    server.StartServer("3000")
}
