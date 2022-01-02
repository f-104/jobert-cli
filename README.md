<div align="center">
  <img alt="Logo goes here" style="width: 25%; height: auto" src="logo.png">
</div>
<h1 align="center">Jobert CLI</h1>
<p align ="center">
  CLI tool for interfacing with the Jobert API
</p>
<div align="center">
  <a href="LICENSE"><img alt="License badge" src="https://img.shields.io/github/license/f-104/jobert-cli?color=blue"></a>
  <a href="https://github.com/f-104/jobert-api/issues"><img alt="GitHub issues" src="https://img.shields.io/github/issues/f-104/jobert-cli?color=blue"></a>
  <img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/f-104/jobert-cli?color=blue">
  <a href="https://github.com/f-104/jobert-cli/releases"><img alt="GitHub Release" src="https://img.shields.io/github/v/release/f-104/jobert-cli?color=blue&include_prereleases"></a>
  <img alt="Go version" src="https://img.shields.io/github/go-mod/go-version/f-104/jobert-cli?color=blue">
</div>


### Table of Contents
- [Summary](#Summary)
- [Installation](#Installation)
- [Usage](#Usage)
- [License](#License)

## Summary
The Jobert CLI is a tool built to interface with a local deployment of the [Jobert API](https://github.com/jobert-app/jobert-api) and facilitate the rapid submission of job applications. Available commands include getting jobs or queries stored in the database, opening jobs from the database, and creating or deleting queries. Jobs can be filtered by their query of origin.

This is an intermediary tool. In the future, Jobert will be offered as a more accessible web application.

## Installation
The recommended approach is to download the appropriate release for your operating system and use the provided executable file. Otherwise, the program can be compiled from the source code provided in this repository. See `go.mod` for dependency information.

## Usage  
The Jobert CLI tool is designed to exchange data with the Jobert API. As there is currently no publicly hosted API deployment, the API is assumed to be accessible at `http://localhost:8080`.

This tool offers a simplified approach to adding new queries and managing data already in the database:
```
Usage:
  jobert [command] [subcommand] [flag]

Available Commands:
  del         Delete queries from database
  get         Get jobs or queries from database
  help        Help about any command
  new         Create a new query
  open        Open jobs in browser

Available Subcommands:
  query       Apply command to queries in the API database
  job         Apply command to jobs in the API database

Flags:
  -h, --help      help for jobert
  -q, --qid (int)   query_id to filter jobs (optional)
```

Some commands have additional restrictions. Use `jobert [command] --help` for more information about a command.
## License
Jobert CLI is licensed under the GPL-3.0 License.