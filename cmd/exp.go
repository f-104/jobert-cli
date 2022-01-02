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
		var url string

		if qid == -1 {
			url = fmt.Sprintf("http://localhost:8080/%s", args[0])
		} else if args[0] == "query" {
			url = fmt.Sprintf("http://localhost:8080/%s/%d", args[0], qid)
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

		//dt := time.Now().Format("01-02-2006")
		if args[0] == "query" {
			var allQueries []Query
			err = json.Unmarshal(body, &allQueries)
			if err != nil {
				log.Fatalln(err)
			}

			//aqt := fmt.Sprintf("./allQueries_%s.csv", dt)
			//f, err := os.Create(aqt)
			//if err != nil {
			//	log.Fatalln(err)
			//}
			// TODO allQueries to CSV
		} else {
			var allJobs []Job
			err = json.Unmarshal(body, &allJobs)
			if err != nil {
				log.Fatalln(err)
			}

			// TODO allQueries to CSV
		}
	},
}

func init() {
	rootCmd.AddCommand(expCmd)
}
