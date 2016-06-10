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
	Name        string `json:"task,omitempty"`
	Description string `json:"description,omitempty"`
	Run         string `json:"run,omitempty"`
}

var tm = map[string]Task{}

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

	// TODO create the map while reading the Dogfile
	for _, t := range tl {
		tm[t.Name] = t
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
		task := tm[arg].Name
		run := tm[arg].Run
		duration := dog.ExecTask(task, []byte(run))
		fmt.Println(duration.Seconds())
	}
}
