package enums

import (
	"fmt"
	_ "os"

	embedfiles "github.com/HarudaySharma/MyAnimeList-CLI/internal/shared/embedFiles"
)

var (
    ApiUrl = ""
)

func init() {
    ApiUrl = fmt.Sprintf("http://localhost:%s/api", embedfiles.DaemonPort)

}

