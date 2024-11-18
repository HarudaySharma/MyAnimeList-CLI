package config

import (
	embedfiles "github.com/HarudaySharma/MyAnimeList-CLI/internal/shared/embedFiles"
)

var C = struct {
	MalApiUrl string
	MalClientId   string
}{
	MalApiUrl: "",
	MalClientId:   "",
}

func init() {
	C.MalClientId = embedfiles.MalClientId
	C.MalApiUrl = embedfiles.MalApiUrl
}
