package daemon

import (
	"os"
	"strings"
)

func IsRunning() bool {
    if !checkFileExists(status_file) {
        return false
    }

	data, err := os.ReadFile(status_file)
	if err != nil {
		panic(err)
	}

	state := strings.Split(string(data), "\n")[0]

	if strings.TrimSpace(strings.ToLower(state)) == "running" {
		return true
	}

	return false
}
