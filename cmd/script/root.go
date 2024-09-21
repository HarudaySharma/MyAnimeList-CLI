package script

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	t "github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mal-cli",
	Short: "Search about anime from terminal",
	Long:  `Access MyAnimList api from terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello to the start")
		// JSON data array
		//req := utils.CreateHttpRequest("GET", "http://localhost:3000/anime/q=\"naruto\"");
		res, err := http.Get(`http://localhost:3000/api/anime-list?q=naruto&fields=`)
		if err != nil {
			fmt.Println(err)
			return
		}

		var animeList t.NativeAnimeList
		if err := json.NewDecoder(res.Body).Decode(&animeList); err != nil {
			fmt.Println("herre")
			fmt.Println(err)
			return

		}
        // show the list of the all the anime found
        for _, v := range(animeList.Data) {
            fmt.Println(v.Title)
        }

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
}
