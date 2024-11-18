package main

import (
	"fmt"
	"strconv"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script"
	"github.com/HarudaySharma/MyAnimeList-CLI/internal/shared/daemon"
	embedfiles "github.com/HarudaySharma/MyAnimeList-CLI/internal/shared/embedFiles"
)

var (
	DaemonPort int64 = 0
)

func main() {
	// start the daemon (mal_cli_api) if not running already
	script.Execute()
    return
	if !daemon.IsRunning() {
		fmt.Println("DAEMON NOT RUNNING")
		fmt.Println("STARTING DAEMON....")
		daemon.StartDaemon(DaemonPort)
	}

    if daemon.IsRunning() {
        fmt.Println("___DAEMON RUNNING!!___")
    }

}

func init() {
    var err error
    DaemonPort, err = strconv.ParseInt(embedfiles.DaemonPort, 10, 64)
    if err != nil {
		panic(fmt.Sprintf("error parsing dameon port %v", err))
    }
}
