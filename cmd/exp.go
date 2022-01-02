package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// expCmd represents the exp command
var expCmd = &cobra.Command{
	Use:   "exp",
	Short: "Export queries or jobs to csv file",
	Long: `Export list of all queries or jobs to a csv file in the current directory.
Optionally use the -q/--qid flag to filter jobs by query_id.

Proper usage: jobert exp [query/job] -q #`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("argument 'query' or 'job' required")
		}
		if args[0] == "query" || args[0] == "job" {
			return nil
		}
		return fmt.Errorf("invalid command 'jobert exp %s'", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exp called")
		// TODO
	},
}

func init() {
	rootCmd.AddCommand(expCmd)
}

// TODO
// 1. Do get.go
// 2. Store all data in list of structs
// 3. Create CSV file
// 4. Print list to CSV file
