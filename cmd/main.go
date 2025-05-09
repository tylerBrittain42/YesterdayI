package main

import (
	"flag"
	"fmt"

	"github.com/tylerBrittain42/YesterdayI/internal/add"
	"github.com/tylerBrittain42/YesterdayI/internal/config"
)

func main() {
	var conf config.Config
	flag.BoolVar(&conf.Command, "add", true, "help for task")
	flag.StringVar(&conf.Task, "t", "", "help for task")
	flag.StringVar(&conf.JiraTicket, "j", "", "help for jira")
	flag.Parse()

	if conf.Command {
		add.Add(&conf)

	} else {
		fmt.Println("calling view here until I research more about subcommands")
	}

}
