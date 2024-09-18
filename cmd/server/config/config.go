package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var C = struct {
	MAL_API_URL string
	MAL_CLIENT_ID   string
}{
	MAL_API_URL: "",
	MAL_CLIENT_ID:   "",
}

func init() {
	//err := godotenv.Load("./../../.env") // for testing
    err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading .env file")
	}

	MAL_CLIENT_ID := os.Getenv("MAL_CLIENT_ID")
	/* if !ok {
	    log.Panic("Please add the CLIENT_ID in .env file")
	} */
	C.MAL_CLIENT_ID = MAL_CLIENT_ID

	MAL_API_URL := os.Getenv("MAL_API_URL")
	/* if !ok {
	    log.Panic("Please add the MAL_API_URL in .env file")
	} */

	C.MAL_API_URL = MAL_API_URL
}
