package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func StopDaemon() {
	if !IsRunning() {
		fmt.Println("Daemon is Not Running")
		return
	}

	data, err := os.ReadFile(status_file)
	if err != nil {
		panic(err)
	}

	pid := strings.Split(strings.Split(string(data), "\n")[1], "=")[1]

	if err := exec.Command("kill", "-9", pid).Start(); err != nil {
		fmt.Println(err)
		fmt.Printf("error stopping the daemon on [pid: %s]\n", pid)
		return
	}

	fmt.Printf("Daemon Stopped [pid = %s]\n", pid)

	// TODO: think about when to write the status file
	//      before or after the proccess is stopped??.

	if err := writeToStatusFile(fmt.Sprintf("STOPPED\n")); err != nil {
		if !checkFileExists(status_file) {
			return
		}

		fmt.Println(err) // this should not happen after the proccess is stopped
	}

}
