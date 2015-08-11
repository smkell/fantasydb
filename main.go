package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strconv"
	"time"
)

func main() {

	app := cli.NewApp()
	app.Name = "fantasydb"
	app.Usage = "query fantasy stats to make informed decisions"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "db",
			Value:  "mongodb://localhost:27017 /fantasydb",
			Usage:  "database backend to connect to",
			EnvVar: "FANTASYDB_DB",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "import",
			Usage: "imports statistics from a url into the database",
			Description: `The import command is used to import statistics into the database. These statistics may be 
projections for a particular year, or they might be actual statistics from a previous year.

The command supports inputs in the following formats: csv, json.

If any arguments are given then they are interpretted as URLs or as file paths for the import.
Otherwise the input will be read from STDIN.`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "type, t",
					Value: "projection",
					Usage: "the type of stats being imported. Must be one of 'projection', or 'actual' (optional)",
				},
				cli.StringFlag{
					Name:  "year, y",
					Value: "[current]",
					Usage: "the year the stats are for",
				},
			},
			Action: importStatsCommand,
		},
	}
	app.Action = func(c *cli.Context) {
		fmt.Println("Fantasy Database v0.1.0")
		fmt.Println(c.Args())
	}

	app.Run(os.Args)
}

func importStatsCommand(c *cli.Context) {
	fmt.Println("Importing stats.")
	statType := c.String("type")
	year := c.String("year")
	args := c.Args()

	if statType != "projection" && statType != "actual" {
		fmt.Printf("Error: Invalid stat type '%s'. Stat type must be 'projection, or 'actual'.", statType)
		os.Exit(1)
	}

	var yearInt int64
	var err error
	if year == "[current]" {
		now := time.Now()
		yearInt = int64(now.Year())
	} else {
		yearInt, err = strconv.ParseInt(year, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Processing stats with type:", statType)
	fmt.Println("Processing stats for year:", yearInt)
	fmt.Println("Importing from sources:", args)
}
