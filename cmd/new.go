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
		var (
			newTerm   string
			newCity   string
			newState  string
			newRadius string
		)

		fmt.Println("Queries contain a search term, city, state, and radius.")
		fmt.Println("Provide state as two-letter abbreviation.")
		fmt.Println("Choose radius from [0, 5, 10, 15, 25, 50, 100].")

		fmt.Print("Enter search term: ")
		fmt.Scanln(&newTerm)
		fmt.Print("Enter city: ")
		fmt.Scanln(&newCity)
		fmt.Print("Enter state: ")
		fmt.Scanln(&newState)
		fmt.Print("Enter radius: ")
		fmt.Scanln(&newRadius)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

// TODO
// 1. [x] Get user input for query data (term, city, state, radius)
// 2. [X] Validate user input
// 3. [ ] Create and marshal new Query struct
// 4. [ ] Send POST request
// 5. [ ] Print statement based on response
