package script

import (
	"fmt"
	"os"
	"strings"

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


func init () {
	// option: --list-size
	rootCmd.PersistentFlags().IntP("list-size", "l", 10, strings.TrimSpace(fmt.Sprintf(`
        Specify length of anime list { 1 - 100 }
        `,
	)))
}
