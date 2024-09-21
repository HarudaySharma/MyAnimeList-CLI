package script

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search anime by title",
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		url := fmt.Sprintf("%s/anime-list?q=%s",
			enums.API_URL,
			query,
		)

		res, err := http.Get(url)
		if err != nil {
			panic("error getting anime list")
		}

		var list types.NativeAnimeList
		if err := json.NewDecoder(res.Body).Decode(&list); err != nil {
			fmt.Printf("Json parsing error of anime-list \n %v", err)
			return
		}
		// show the list with fzf
		titleList := strings.Builder{}
		titleMap := make(map[string]int, 0)
		for _, val := range list.Data {
			titleList.WriteString(val.Title + "\n")
			titleMap[val.Title+"\n"] = val.ID
		}

		fzf := exec.Command("fzf")
		fzf.Stdin = strings.NewReader(titleList.String())

		selectedTitle, err := fzf.Output()
		if err != nil {
			fmt.Printf("error using fzf \n %v", err)
			return
		}

		// get the anime information
		url = fmt.Sprintf("%s/anime/%d?detail_type=%s",
			enums.API_URL,
			titleMap[string(selectedTitle)],
			"basic",
		)

		res, err = http.Get(url)
		if err != nil {
			panic("error getting anime list")
		}

		var animeDetails types.NativeAnimeDetails_Basic
		if err := json.NewDecoder(res.Body).Decode(&animeDetails); err != nil {
			fmt.Printf("Json parsing error of animeDetails \n %v", err)
			return
		}

		// FIX:  textView being inferred as *tview.Box
		// therefore, not getting the SetText() function access
		// and also the tview application is all empty
		textView := tview.NewTextView().
			SetTitle(string(selectedTitle)).
            SetText(animeDetails.Synopsis).
			SetTitleColor(tcell.ColorGreen).
			SetBackgroundColor(tcell.ColorDefault)

		// Create a new application
		// Set the root and run the application
		if err := tview.NewApplication().SetRoot(textView, true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
