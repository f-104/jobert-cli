package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open jobs in browser",
	Long: `Open jobs currently in database using default browser.
Optionally use the -q/--qid flag to filter jobs by query_id.

Proper usage: jobert open -q #`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			url    string
			reader *bufio.Reader = bufio.NewReader(os.Stdin)
		)
		if qid == -1 {
			url = "http://localhost:8080/job"
		} else {
			url = fmt.Sprintf("http://localhost:8080/job?query_id=%d", qid)
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

		var allJobs []Job
		err = json.Unmarshal(body, &allJobs)
		if err != nil {
			log.Fatalln(err)
		}

		l := len(allJobs)

		if l == 0 {
			log.Fatalln("no jobs found")
		}
		fmt.Printf("%d jobs found. Are you sure you want to open? [y/N]", l)
		ans, _ := reader.ReadString('\n')
		winAns := strings.Replace(ans, "\r\n", "", -1)
		if winAns != "y" {
			os.Exit(0)
		}

		for i := 0; i < l; i++ {
			url := allJobs[i].Href
			err = browser.OpenURL(url)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
