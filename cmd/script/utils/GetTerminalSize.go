package utils

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetTerminalSize() (int, int) {
	fd := int(os.Stdout.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return 80, 24
	}

	return width, height
}
