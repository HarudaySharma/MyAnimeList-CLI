package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
)


func ImageToBase64(img image.Image) (string) {
    buf := new(bytes.Buffer)
    if err := jpeg.Encode(buf, img, nil); err != nil {
        fmt.Println("failed to encode image")
        return ""
    }

    return base64.StdEncoding.EncodeToString(buf.Bytes())
}


