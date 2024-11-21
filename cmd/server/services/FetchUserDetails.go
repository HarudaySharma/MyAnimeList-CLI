package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	c "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/config"
	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	u "github.com/HarudaySharma/MyAnimeList-CLI/pkg/utils"
)

func FetchUserDetails() *t.NativeUserDetails {
	client := http.Client{}

	url := fmt.Sprintf("%s/users/@me",
		c.C.MalApiUrl,
	)
	log.Println(url)
	req := u.CreateUserHttpRequest("GET", url)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR in FetchUserDetails fetching from MAL API \n %v", err)
		log.Fatal(err)
		return &t.NativeUserDetails{}
	}

	var data t.MALUserDetails
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatalf("ERROR in FetchAnimeDetails decoding json body \n %v", err)
	}
	defer res.Body.Close()

	if CheckInvalidAccessToken(data) {
		log.Printf("access_token invalid or expired probably")
		// try generating new access_token
		if err := UpdateAccessToken(); err != nil {
			// try generating new refresh_token
			if err := UpdateRefreshToken(); err != nil {
				return &t.NativeUserDetails{} // return nice error
			}
		}
		return FetchUserDetails()
	}

	return convertToNativeUserDetails(&data)
}

func convertToNativeUserDetails(data *t.MALUserDetails) *t.NativeUserDetails {
	convertedData := t.NativeUserDetails{}

	convertedData.Id = data.Id
	convertedData.Name = data.Name
	convertedData.Location = data.Location
	convertedData.AnimeStatistics = data.AnimeStatistics
	convertedData.JoinedAt = data.JoinedAt

	return &convertedData
}

func CheckInvalidAccessToken(v any) bool {
	if d, ok := v.(t.MALUserDetails); ok {
		if d.Error == "invalid_token" {
			return true
		}
	}
	if errMap, ok := v.(map[string]interface{}); ok {
		if error, ok := errMap["error"].(string); ok && error == "invalid_token" {
			return true
		}
	}
	return false
}
