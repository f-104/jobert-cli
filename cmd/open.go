package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open jobs in browser",
	Long: `Open jobs currently in database using default browser.
Optionally use the -q/--qid flag to filter jobs by query_id.

Proper usage: jobert open -q #`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("open called")
		// TODO
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
