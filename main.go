package main

import "github.com/jobert-app/jobert-cli/cmd"

// Each command is stored in /cmd as an individual file.
// Command functionality is enabled all at once with cmd.Execute()
func main() {
	cmd.Execute()
}
