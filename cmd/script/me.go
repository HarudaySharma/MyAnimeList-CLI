package script

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "get user specific data",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(fmt.Sprintf("http://localhost:42069/api/user"))
		if err != nil {
			fmt.Println(err)
			return
		}

		var data any
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			fmt.Println(err)
		}

		jsonData, _ := json.MarshalIndent(data, "\t", "")
		fmt.Println(string(jsonData))
	}}
