package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
)

func FetchImage(imageUrl string) image.Image {
    resp, err := http.Get(imageUrl)
    if err != nil {
        fmt.Println("error fetching photo")
        return nil
    }
    defer resp.Body.Close()

    imgData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err, "failed reading request body")
        return nil
    }

    img, err := jpeg.Decode(bytes.NewReader(imgData))
    if err != nil {
        fmt.Println(err, "failed to decode JPEG")
        return nil

    }

    return img
}
