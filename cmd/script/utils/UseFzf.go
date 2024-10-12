package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func UseFzf(input []string) (string, error) {
	fzf := exec.Command("fzf", "--no-sort", "--cycle", "--ansi", "+m")
	fzf.Stdin = strings.NewReader(strings.Join(input, "\n"))

	output, err := fzf.Output()
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

    return string(output), nil
}
