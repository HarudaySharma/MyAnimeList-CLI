package utils

import (
	"bytes"
	"fmt"
	"image"
	"os"
)

// FetchImage returns the image and its mimetype
func FetchImage(imageId int, imageUrl string) (image.Image, string)   {
    //TODO: add caching here too

    filePath := imageDir + "/id/" + fmt.Sprintf("%d", imageId)
    if err := DownloadImage(imageUrl,  filePath); err != nil {
        fmt.Println("error downloading image")
        return nil, ""
    }

    imgData, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println(err, "failed reading request body")
        return nil, ""
    }

    img, format, err := image.Decode(bytes.NewReader(imgData))
    if err != nil {
        fmt.Println(err, "failed to decode image")
    }

    return img, format;
}
