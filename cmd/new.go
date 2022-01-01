package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new query",
	Long: `Add a new query to the API database by providing information as requested:
term: [term]
city: [city]
state: [state] (two-letter abbreviation)
radius: [radius] (distance from city to search)

Proper usage: 'jobert new query'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("argument 'query' required")
		}
		if args[0] == "query" {
			return nil
		}
		return fmt.Errorf("invalid command 'jobert new %s'", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
		// TODO
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
