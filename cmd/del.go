package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete queries or jobs from database",
	Long: `Delete queries or jobs from API database.
Optionally use the -q/--qid flag to filter by query_id.

WARNING: If -q not specified, will delete all entries.

Proper usage: jobert del [query/job] -q #`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("argument 'query' or 'job' required")
		}
		if args[0] == "query" || args[0] == "job" {
			return nil
		}
		return fmt.Errorf("invalid command 'jobert del %s'", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("del called")
		// TODO
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
