package utils

import "os"

func DeleteFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
        return err
	}

    return nil
}
