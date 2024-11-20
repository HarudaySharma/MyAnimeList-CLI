package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadConfigFile(key string) string {
    filePath := os.ExpandEnv("/$HOME/.config/mal-cli/config")
    data, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("no config file found")
        return ""
    }

    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        words := strings.Split(line, "=")
        if len(words) <= 1 {
            continue
        }

        k := words[0]
        v := words[1]
        if k == key {
            return v
        }
    }

    return ""
}
