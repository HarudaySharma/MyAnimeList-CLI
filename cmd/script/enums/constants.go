package enums

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
    API_URL = ""
)

func init() {
    err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	DAEMON_PORT := os.Getenv("DAEMON_PORT")
    API_URL = fmt.Sprintf("http://localhost:%s/api", DAEMON_PORT)
}

