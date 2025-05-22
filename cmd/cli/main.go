package main

import (
	"flag"
	"fmt"

	"github.com/tylerBrittain42/YesterdayI/internal/config"
	"github.com/tylerBrittain42/YesterdayI/internal/task"
)

const fileName = "taskLog.json"

func main() {

	var conf config.Config
	flag.BoolVar(&conf.Command, "add", true, "help for task")
	flag.StringVar(&conf.Content, "t", "", "help for task")
	flag.StringVar(&conf.JiraTicket, "j", "", "help for jira")
	flag.Parse()

	if conf.Command {
		err := task.AddTask(fileName, &conf)
		if err != nil {
			fmt.Printf("unable to add task: %v\n", err)
		} else {
			fmt.Println("task added")
		}

	} else {
		fmt.Println("calling view here until I research more about subcommands")
	}

}
