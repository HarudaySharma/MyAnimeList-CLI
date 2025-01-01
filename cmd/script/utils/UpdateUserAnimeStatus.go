package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
)

type UpdateUserAnimeStatusParams struct {
	AnimeId     int
	AnimeStatus *types.NativeUserAnimeStatus // this will be updated
}

func UpdateUserAnimeStatus(p UpdateUserAnimeStatusParams) error {
	url := fmt.Sprintf("%s/user/anime/%d/my_list_status",
		enums.ApiUrl,
		p.AnimeId,
	)

	jsonData, err := json.Marshal(p.AnimeStatus)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%v\n****ERROR UPDATING THE USER ANIME STATUS FROM SERVER****", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 && res.StatusCode < 600 {
		var v any
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			return fmt.Errorf("Json parsing error of anime-list \n %v", err)
		}
		jsonData, _ := json.MarshalIndent(v, "\t", " ")

		return fmt.Errorf("%s", jsonData)
	}

	if err := json.NewDecoder(res.Body).Decode(p.AnimeStatus); err != nil {
		return fmt.Errorf("Json parsing error of anime-list \n %v", err)
	}

	return nil

}
