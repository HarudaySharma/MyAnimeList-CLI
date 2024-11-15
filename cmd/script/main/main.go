package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script"
	"github.com/HarudaySharma/MyAnimeList-CLI/internal/shared/daemon"
	"github.com/joho/godotenv"
)

var (
	DAEMON_PORT int64 = 0
)

func main() {
	// start the daemon (mal_cli_api) if not running already
	script.Execute()
    return
	if !daemon.IsRunning() {
		fmt.Println("DAEMON NOT RUNNING")
		fmt.Println("STARTING DAEMON....")
		daemon.StartDaemon(DAEMON_PORT)
	}

    if daemon.IsRunning() {
        fmt.Println("___DAEMON RUNNING!!___")
    }

}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

    DAEMON_PORT, err = strconv.ParseInt(os.Getenv("DAEMON_PORT"), 10, 64)
    if err != nil {
		panic("error parsing DAEMON_PORT")
    }
}
