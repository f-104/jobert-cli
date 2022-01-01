package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get jobs or queries from database",
	Long: `Get jobs or queries from API database.
Optionally use the -q/--qid flag to filter by query_id.

Proper usage: jobert get [query/job] -q #`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("argument 'query' or 'job' required")
		}
		if args[0] == "query" || args[0] == "job" {
			return nil
		}
		return fmt.Errorf("invalid command 'jobert get %s'", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
