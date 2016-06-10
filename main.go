package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghodss/yaml"

	"github.com/xsb/dog/dog"
)

type TaskList []Task

type Task struct {
	Task        string `json:"task,omitempty"`
	Description string `json:"description,omitempty"`
	Run         string `json:"run,omitempty"`
}

// var tl = TaskList{
// 	Task{
// 		Task:        "hello",
// 		Description: "Say Hello!",
// 		Run:         "echo hello world",
// 	},
// 	Task{
// 		Task:        "bye",
// 		Description: "Good Bye!",
// 		Run:         "echo bye cruel world",
// 	},
// }

var tm = map[string]Task{
	"hello": Task{
		Description: "Say Hello!",
		Run:         "echo hello world",
	},
	"bye": Task{
		Description: "Good Bye!",
		Run:         "echo bye cruel world",
	},
	"find": Task{
		Description: "List all files in $HOME directory",
		Run:         "find /home/xavi",
	},
}

func loadDogFile() (tl TaskList, err error) {
	var dat []byte
	dat, err = ioutil.ReadFile("dogfile/testdata/Dogfile-basic.yml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(dat, &tl)
	if err != nil {
		return
	}
	return
}

func main() {
	tl, err := loadDogFile()
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(tl)

	arg := os.Args[1]
	if arg == "list" || arg == "help" {
		for k, t := range tm {
			fmt.Printf("%s\t%s\n", k, t.Description)
		}
	} else {
		// TODO check that task exists
		task := tm[arg].Task
		run := tm[arg].Run
		duration := dog.ExecTask(task, []byte(run))
		fmt.Println(duration.Seconds())
	}
}
