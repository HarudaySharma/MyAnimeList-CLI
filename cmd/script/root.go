package script

import (
	"fmt"
	"os"
	"strings"

	usercommands "github.com/HarudaySharma/MyAnimeList-CLI/cmd/script/user-commands"
	es "github.com/HarudaySharma/MyAnimeList-CLI/cmd/server/enums"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mal-cli",
	Short: "Search about anime from terminal",
	Long:  `Access MyAnimeList api from terminal`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// option: --list-size
	rootCmd.PersistentFlags().IntP("list-size", "l", 10, strings.TrimSpace(fmt.Sprintf(`
        Specify length of anime list { 1 - 100 }
        `,
	)))

	// option: --details
	availableOptionsStr := strings.Builder{}
	availableOptionsStr.WriteString("\n\t\t")

	everyDetailField := es.EveryDetailField()
	for i := 0; i < len(everyDetailField); i+=2 {
		availableOptionsStr.WriteString(fmt.Sprintf("%d => %s", i, everyDetailField[i]))
		availableOptionsStr.WriteString("\t\t\t\t\t")
		if i+1 < len(everyDetailField) {
			availableOptionsStr.WriteString(fmt.Sprintf("%d => %s", i + 1, everyDetailField[i + 1]))
		}
		availableOptionsStr.WriteString("\n\t\t")
	}

	rootCmd.PersistentFlags().IntSliceP("details", "d", []int{}, strings.TrimSpace(fmt.Sprintf(`
        Specify which anime detail you want

            Available Options: %s
        `,
		availableOptionsStr.String(),
	)))

    rootCmd.AddCommand(usercommands.MeCmd)
}
