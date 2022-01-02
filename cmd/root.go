package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// optionally filter jobs/queries by query_id
var qid int

// structs correspond to tables in API database
type Query struct {
	City   string `json:"city"`
	Id     int    `json:"id"`
	Radius string `json:"radius"`
	State  string `json:"state"`
	Term   string `json:"term"`
}

type Job struct {
	Company  string `json:"company"`
	Href     string `json:"href"`
	Id       int    `json:"id"`
	Location string `json:"location"`
	Query_id int    `json:"query_id"`
	Title    string `json:"title"`
}

// base command
var rootCmd = &cobra.Command{
	Use:   "jobert",
	Short: "CLI tool to interface with Jobert API",
	Long: `The Jobert CLI tool is designed to exchange data with the Jobert API.
As there is currently no publicly hosted API deployment, the API is assumed
to be accessible at https://localhost:8080.

This tool offers a simplified approach to adding new queries and managing
data already in the database.`,
}

// Adds all child commands to the root command and sets flags. Called by main.main()
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&qid, "qid", "q", -1, "query_id to filter jobs")
}
