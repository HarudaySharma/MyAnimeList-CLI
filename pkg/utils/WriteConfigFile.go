package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteConfigFile(key string, value string) error {
	var file *os.File
	var err error

	filePath := os.ExpandEnv("$HOME/.config/mal-cli/config")
	if !checkFileExists(filePath) {
		dir := filePath[:strings.LastIndex(filePath, "/")]
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directories: %v", err)
		}
		file, err = os.Create(filePath)
	} else {
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	}

	defer file.Close()

	if err != nil {
		return fmt.Errorf("error opening config file at %s: %v", filePath, err)
	}

	// Read the file's current content
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading config file at %s: %v", filePath, err)
	}

	updatedData := strings.Builder{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			continue // Skip the old key-value pair
		}
		if len(line) > 0 {
			updatedData.WriteString(line)
			updatedData.WriteByte('\n')
		}
	}

	// Add the new key-value pair
	updatedData.WriteString(fmt.Sprintf("%s=%s\n", key, value))

	// Truncate and write the updated content
	err = file.Truncate(0)
	if err != nil {
		return fmt.Errorf("error truncating config file at %s: %v", filePath, err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("error seeking config file at %s: %v", filePath, err)
	}
	_, err = file.WriteString(updatedData.String())
	if err != nil {
		return fmt.Errorf("error writing to config file at %s: %v", filePath, err)
	}

	return nil
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

