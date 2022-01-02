package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get jobs or queries from database",
	Long: `Get jobs or queries from API database.
Optionally use the -q/--qid flag to filter jobs by query_id.

Proper usage: jobert get query
			  jobert get job -q #`,
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
		var url string

		if qid == -1 {
			url = fmt.Sprintf("http://localhost:8080/%s", args[0])
		} else if args[0] == "query" {
			log.Fatalln("-q/--qid only supported for filtering jobs")
		} else {
			url = fmt.Sprintf("http://localhost:8080/%s?query_id=%d", args[0], qid)
		}

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(resp.Status)

		if args[0] == "query" {
			var allQueries []Query
			err = json.Unmarshal(body, &allQueries)
			if err != nil {
				log.Fatalln(err)
			}

			for i := 0; i < len(allQueries); i++ {
				fmt.Println("id:", allQueries[i].Id, "\t", "radius:", allQueries[i].Radius)
				fmt.Println("term:", allQueries[i].Term)
				fmt.Println("location:", allQueries[i].City+", "+allQueries[i].State)
				fmt.Println("-----")
			}

			if len(allQueries) == 1 {
				fmt.Println(len(allQueries), "query")
			} else {
				fmt.Println(len(allQueries), "queries")
			}
		} else {
			var allJobs []Job
			err = json.Unmarshal(body, &allJobs)
			if err != nil {
				log.Fatalln()
			}

			for i := 0; i < len(allJobs); i++ {
				fmt.Println("id:", allJobs[i].Id, "\t", "query_id:", allJobs[i].Query_id)
				fmt.Println("title:", allJobs[i].Title)
				fmt.Println("company:", allJobs[i].Company)
				fmt.Println("location:", allJobs[i].Location)
				fmt.Println("-----")
			}

			if len(allJobs) == 1 {
				fmt.Println(len(allJobs), "job")
			} else {
				fmt.Println(len(allJobs), "jobs")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
