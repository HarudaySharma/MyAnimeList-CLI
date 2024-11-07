package script

import (
	"fmt"
	"os"
	"strings"

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
	for i, option := range es.EveryDetailField() {
		availableOptionsStr.WriteString(fmt.Sprintf("%d => %s", i, option))
		availableOptionsStr.WriteString("\n\t\t")
	}

	rootCmd.PersistentFlags().IntSliceP("details", "d", []int{}, strings.TrimSpace(fmt.Sprintf(`
        Specify which anime detail you want

            Available Options: %s
        `,
		availableOptionsStr.String(),
	)))

}
