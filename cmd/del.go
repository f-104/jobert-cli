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

Proper usage: jobert del query -q #`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("argument 'query' or 'job' required")
		}
		if args[0] == "query" {
			return nil
		}
		return fmt.Errorf("invalid command 'jobert del %s'", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		if qid == -1 {
			log.Fatal("-q/--qid is required.")
		}

		client := &http.Client{}

		url := fmt.Sprintf("http://localhost:8080/query?id=%d", qid)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		//
		fmt.Println(resp.Status)
		//fmt.Println(resp.Header)
		fmt.Sprintln(string(b))
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
