package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tylerBrittain42/YesterdayI/internal/config"
	"github.com/tylerBrittain42/YesterdayI/internal/task"
)

func main() {

	fileName := ".yi.json"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileName = filepath.Join(homeDir, fileName)

	var conf config.Config

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCmd.StringVar(&conf.Content, "t", "", "required task")
	addCmd.StringVar(&conf.JiraTicket, "j", "", "optional jira ticket")

	viewCmd := flag.NewFlagSet("view", flag.ExitOnError)
	viewCmd.StringVar(&conf.StartTime, "start", "", "specify a starting date in 00/00 format")
	viewCmd.StringVar(&conf.EndTime, "end", "", "specify an inclusive ending time in 00/00 format")
	viewCmd.StringVar(&conf.SpecificTime, "date", "", "specify a day in 00/00 format")
	viewCmd.BoolVar(&conf.IsJson, "json", false, "will output in json format")

	if len(os.Args) < 2 {
		fmt.Println("No sub command given")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		err := task.AddTask(fileName, &conf)
		if err != nil {
			fmt.Printf("Unable to add task: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("task added")

	case "view":
		viewCmd.Parse(os.Args[2:])
		err := task.View(fileName, &conf)
		if err != nil {
			fmt.Printf("Unable to view tasks: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Println("unexpected sub command given")
		os.Exit(1)
	}

}
