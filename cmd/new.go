package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

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

radius must be chosen from [0, 5, 10, 15, 25, 50, 100].

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
		// id is autoincremented and should not be included in a POST request
		type postQuery struct {
			City   string `json:"city"`
			Radius string `json:"radius"`
			State  string `json:"state"`
			Term   string `json:"term"`
		}

		var (
			reader    *bufio.Reader = bufio.NewReader(os.Stdin)
			newTerm   string
			newCity   string
			newState  string
			newRadius string
			newQuery  postQuery
		)

		fmt.Println("Queries contain a search term, city, state, and radius.")
		fmt.Println("Provide state as two-letter abbreviation.")
		fmt.Println("Choose radius from [0, 5, 10, 15, 25, 50, 100].")

		fmt.Print("Enter search term: ")
		newTerm, _ = reader.ReadString('\n')
		fmt.Print("Enter city: ")
		newCity, _ = reader.ReadString('\n')
		fmt.Print("Enter state: ")
		newState, _ = reader.ReadString('\n')
		fmt.Print("Enter radius: ")
		newRadius, _ = reader.ReadString('\n')

		// only need to replace \r\n due to Windows development. Will be changed in a future commit
		newQuery.Term = strings.Replace(newTerm, "\r\n", "", -1)
		newQuery.City = strings.Replace(newCity, "\r\n", "", -1)
		newQuery.State = strings.Replace(newState, "\r\n", "", -1)
		newQuery.Radius = strings.Replace(newRadius, "\r\n", "", -1)

		mQuery, err := json.Marshal(newQuery)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Post("http://localhost:8080/query", "application/json", bytes.NewBuffer(mQuery))
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
