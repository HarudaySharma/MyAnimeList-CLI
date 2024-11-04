package daemon

import (
	"fmt"
	"os"
	"os/exec"
)


func StartDaemon(port int64) {
	daemon := os.ExpandEnv("$HOME/projects/MyAnimeList-CLI/bin/main.bin")
	cmd := exec.Command(daemon, fmt.Sprintf("%d", port))

	err := cmd.Start()
	if err != nil {
		fmt.Println("FAILED TO START THE DAEMON")
		panic(err)
	}

    pid := cmd.Process.Pid
    fmt.Printf("Process started with PID %d at PORT :%d\n", pid, port)


	writeToStatusFile(fmt.Sprintf("RUNNING\nPID=%d\n", pid))
}
