package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete queries from database",
	Long: `Delete queries from API database.
Use the -q/--qid flag to filter by query_id (required).

Deleting all entries is not supported by the API.

Jobs older than one day should be automatically deleted on the API backend.

Proper usage: jobert del query -q #`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("argument 'query' required")
		}
		if args[0] == "query" {
			return nil
		}
		return fmt.Errorf("invalid command 'jobert del %s'", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		if qid == -1 {
			log.Fatalln("-q/--qid is required.")
		}

		client := &http.Client{}

		url := fmt.Sprintf("http://localhost:8080/query/%d", qid)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(resp.Status)
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
