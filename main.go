package main

import (
	"fmt"
	"os"

	"github.com/xsb/dog/dog"
)

var taskList = map[string]dog.Task{
	"hello": {
		Name:        "hello",
		Description: "Say Hello!",
		Time:        false,
		Run:         []byte("echo hello world"),
	},
	"bye": {
		Name:        "bye",
		Description: "Good Bye!",
		Time:        true,
		Run:         []byte("echo bye cruel world"),
	},
	"find": {
		Name:        "find",
		Description: "List all files in $HOME directory",
		Time:        true,
		Run:         []byte("find /home/xavi"),
	},
}

func main() {
	arg := os.Args[1]
	if arg == "list" || arg == "help" {
		for k, t := range taskList {
			fmt.Printf("%s\t%s\n", k, t.Description)
		}
	} else {
		// TODO check that task exists
		task := taskList[arg]
		dog.ExecTask(task)
	}
}
