package utils

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
)

// FetchImage returns the image and its mimetype
func FetchImage(imageUrl string) (image.Image, string)   {
    resp, err := http.Get(imageUrl)
    if err != nil {
        fmt.Println("error fetching photo")
        return nil, ""
    }
    defer resp.Body.Close()

    imgData, err := ioutil.ReadAll(resp.Body)
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
