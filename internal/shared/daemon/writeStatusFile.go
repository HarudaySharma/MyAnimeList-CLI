package daemon

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var status_file = filepath.Join("/tmp", "/mal_cli_daemon_status")

func writeToStatusFile(data string) error {

    var file *os.File;
    var err error

    if !checkFileExists(status_file) {
        file, err = os.Create(status_file)
    } else {
        file, err = os.OpenFile(status_file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    }

    defer file.Close()

    if err != nil {
        return errors.New(fmt.Sprintf("error writing to status file at %s", status_file))
    }

    file.WriteString(data)

    return nil
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

