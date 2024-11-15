package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadImage(url, filePath string) error {
    //fmt.Printf("url: %s\n", url)
    if checkFileExists(filePath) {
        //fmt.Println("file exists")
        return nil
    }

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

    dir := filePath[:len(filePath)-len("/"+filePath[strings.LastIndex(filePath, "/")+1:])]

    err = os.MkdirAll(dir, os.ModePerm)
    if err != nil {
        return fmt.Errorf("failed to create directories: %v", err)
    }

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
    //fmt.Printf("%s created\n", filePath)
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}
