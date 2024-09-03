package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var C = struct {
	MAL_API_URL string
	CLIENT_ID   string
}{
	MAL_API_URL: "",
	CLIENT_ID:   "",
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

    log.Print(os.Getenv("CLIENT_ID"))

	CLIENT_ID := os.Getenv("CLIENT_ID")
	/* if !ok {
	    log.Panic("Please add the CLIENT_ID in .env file")
	} */
	C.CLIENT_ID = CLIENT_ID

	MAL_API_URL := os.Getenv("MAL_API_URL")
	/* if !ok {
	    log.Panic("Please add the MAL_API_URL in .env file")
	} */

	C.MAL_API_URL = MAL_API_URL
}
