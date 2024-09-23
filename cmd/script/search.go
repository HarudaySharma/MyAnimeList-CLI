package script

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/utils"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/HarudaySharma/MyAnimeList-CLI/pkg/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search anime by title",
	Run: func(cmd *cobra.Command, args []string) {
		var query string
		if len(args) == 0 || len(args[0]) < 3 {
			for len(query) < 3 {
				fmt.Print("Enter the anime title [atleast 3 letters word]: ")
				fmt.Scanln(&query)
			}
		} else {
			query = args[0]
		}

		var animeList types.NativeAnimeList
		err := utils.GetAnimeList(&animeList, query, 2, 0, []es.AnimeDetailField{
			es.StartSeason,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		// show the list with fzf
		// TODO: give an option for [next list & previous list]

		titleList := make([]string, 0)
		titleMap := make(map[string]int, 0)
		for _, val := range animeList.Data {
			plainKeyStr := strings.Builder{}
			formattedKeyStr := strings.Builder{}

			formattedKeyStr.WriteString(val.Title + "\t")
			plainKeyStr.WriteString(val.Title + " ")

			startSeason := val.CustomFields["start_season"]
			if seasonData, ok := startSeason.(map[string]interface{}); ok {
				season, _ := seasonData["season"]
				year, _ := seasonData["year"]

				plainKeyStr.WriteString(fmt.Sprintf("[ %v ~ %v ]", year, season))
				formattedKeyStr.WriteString(fmt.Sprintf("\033[35m[ %v ~ %v ]\033[0m", year, season))
			}

			titleMap[plainKeyStr.String()] = val.ID
			titleList = append(titleList, strings.TrimSpace(formattedKeyStr.String()))
		}

		fzf := exec.Command("fzf", "--no-sort", "--cycle", "--ansi", "+m")
		fzf.Stdin = strings.NewReader(strings.Join(titleList, "\n"))

		output, err := fzf.Output()
		if err != nil {
			fmt.Printf("error using fzf \n %v", err)
			return
		}

		selectedTitle := strings.TrimSpace(string(output))
		// Strip ANSI codes from selectedTitle to match the titleMap keys
		cleanTitle := stripAnsi(strings.ReplaceAll(selectedTitle, "\t", " "))

		animeId, found := titleMap[cleanTitle]
		if !found {
			fmt.Println("Selected title not found in the map.")
			return
		}

		var animeDetails types.NativeAnimeDetails
		utils.GetAnimeDetails(&animeDetails, animeId, "custom", []es.AnimeDetailField{
			es.Id, es.Title,
			es.Synopsis,
			es.AlternativeTitles,
			es.Genres,
		})

		alternativeTitles := strings.Builder{}
		alternativeTitles.WriteString("EN:\t")
		alternativeTitles.WriteString(fmt.Sprintln(animeDetails.AlternativeTitles.EN))
		alternativeTitles.WriteString("JA:\t")
		alternativeTitles.WriteString(fmt.Sprintln(animeDetails.AlternativeTitles.JA))

		titleBox := tview.NewTextView().
			SetText(alternativeTitles.String()).
			SetTextAlign(tview.AlignLeft).
			SetWrap(true)

		titleBox.SetBackgroundColor(tcell.ColorDefault).
			SetTitle("Title").
			SetTitleAlign(tview.AlignLeft).
			SetTitleColor(tcell.ColorLightCyan).
			SetBorder(true)

		synopsisBox := tview.NewTextView().
			SetText(string(animeDetails.Synopsis)).
			SetTextAlign(tview.AlignLeft).
			SetWrap(true).
			SetScrollable(true)

		synopsisBox.SetBackgroundColor(tcell.ColorDefault).
			SetTitle("Synopsis").
			SetBorder(true)

		genres := strings.Builder{}
		for i, genre := range animeDetails.Genres {
			if i != 0 {
				genres.WriteString(", ")
			}
			genres.WriteString(genre.Name)
		}

		genresBox := tview.NewTextView().
			SetText(genres.String())

		genresBox.SetBackgroundColor(tcell.ColorDefault).
			SetTitle("Genres").
			SetTitleColor(tcell.ColorGreenYellow).
			SetTitleAlign(tview.AlignLeft).
			SetBorder(true)

		leftBox := tview.NewBox().
			SetTitle("Left (20 cols)").
			SetBackgroundColor(tcell.ColorDefault).
			SetBorder(true)

		rightBox := tview.NewBox().
			SetTitle("Right (20 cols)").
			SetBackgroundColor(tcell.ColorDefault).
			SetBorder(true)

		flex := tview.NewFlex().
			AddItem(leftBox, 20, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(titleBox, 5, 1, false).
				AddItem(synopsisBox, 0, 3, false).
				AddItem(genresBox, 5, 1, false), 0, 2, false).
			AddItem(rightBox, 20, 1, false)

		if err := tview.NewApplication().SetRoot(flex, true).Run(); err != nil {
			panic(err)
		}
	},
}

// Helper function to strip ANSI codes from a string
func stripAnsi(str string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(str, "")
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
