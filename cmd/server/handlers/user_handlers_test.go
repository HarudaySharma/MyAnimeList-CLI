package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
)

func Test_GETUserInfo(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:42069/api/user"))
	if err != nil {
		fmt.Println(err)
        t.FailNow()
        return
	}

    var data any
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
        t.FailNow()
    }

    jsonData, _ := json.MarshalIndent(data, "\t", "")
    fmt.Println(string(jsonData))
}

func Test_UserHandlers(t *testing.T) {
	url := fmt.Sprintf("https://myanimelist.net/v1/oauth2/authorize?response_type=code&client_id=%s&code_challenge=%s&state=Request123&code_challenge_method=plain",
		config.C.MalClientId,
		config.C.MalCodeChallenge,
	)

	fmt.Println("go to this url to authorize application to read/write your mal data")
	fmt.Println()
	fmt.Println(url)
	fmt.Println()

	fmt.Print("Enter the code you recieved in the callback url query paramater: ")

	input := "authCode\n"

	reader := bufio.NewReader(strings.NewReader(input))
	authCode := ""

	for len(authCode) == 0 {
		input, _ := reader.ReadString('\n')
		authCode = strings.TrimSpace(input)
	}

	_, err := http.Get(fmt.Sprintf("http://localhost:42069/api/auth?code=%s", authCode))
	if err != nil {
		fmt.Println(err)
	}
}
