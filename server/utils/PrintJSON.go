package utils

import (
	"encoding/json"
	"log"
)

func PrintJSON(v interface{}) {
    d, err := json.MarshalIndent(v, "", "\t")
    if err != nil {
        log.Fatal("Error marshalling JSON")
        return
    }
    log.Println(string(d))

    return
}
