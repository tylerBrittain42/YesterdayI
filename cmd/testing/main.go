package main

import (
	"fmt"

	"github.com/tylerBrittain42/YesterdayI/internal/task"
)

const fileName = "/home/tyler/Documents/YesterdayI/cmd/testing/taskLog.json"

func main() {
	// testSave()
	testLoad()
}

func testSave() {
	tSlice := task.TaskSlice{}
	tSlice.Add("ai did a thing", "config-2159")
	tSlice.Add("I did something else", "config-2158")

	tSlice.Save(fileName)

}

func testLoad() {

	tSlice, _ := task.Load(fileName)
	fmt.Println(tSlice)
}
