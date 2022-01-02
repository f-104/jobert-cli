package cmd

import (
	"os"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
)

// optionally filter jobs/queries by query_id
var qid int

type Query struct {
	City   string
	Id     int
	Radius string
	State  string
	Term   string
}

type Job struct {
	Company  string
	Href     string
	Id       int
	Location string
	Query_id int
	Title    string
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
