package config

import (
	embedfiles "github.com/HarudaySharma/MyAnimeList-CLI/internal/shared/embedFiles"
)

var C = struct {
	MalApiUrl string
	MalClientId   string
    MalAuthCode string
    MalCodeChallenge string
    MalCodeVerifier string
}{
	MalApiUrl: "",
	MalClientId:   "",
	MalCodeChallenge:   "",
}

func init() {
	C.MalClientId = embedfiles.MalClientId
	C.MalApiUrl = embedfiles.MalApiUrl
    C.MalCodeChallenge = embedfiles.MalCodeChallenge
    C.MalCodeVerifier = embedfiles.MalCodeChallenge

    // NOTE: MalAuthCode will be filled from the file
    // save the access_tokens and refresh_tokens you got by using auth_code
    C.MalAuthCode = ""
}
