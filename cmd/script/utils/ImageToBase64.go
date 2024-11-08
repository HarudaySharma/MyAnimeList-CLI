package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
)


// ImageToBase64 return base64 encoded image string and the image mimetype
func ImageToBase64(img image.Image, mimetype string) (string, string) {
    // Image format supported: jpeg
    buf := new(bytes.Buffer)
    if mimetype == "jpg" || mimetype == "jpeg" {
        if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100}); err != nil {
            fmt.Println("failed to encode image of type: ", mimetype);
            return "", mimetype
        }
    }
    if mimetype == "png" {
        if err := png.Encode(buf, img); err != nil {
            fmt.Println("failed to encode image of type: ", mimetype);
            return "", mimetype
        }
    }

    return base64.StdEncoding.EncodeToString(buf.Bytes()), mimetype

}


