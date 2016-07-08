package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dogtools/dog/execute"
	"github.com/dogtools/dog/parser"
	"github.com/joho/godotenv"
)

const version = "0.0"

func main() {
	// if .env file exists (in same dir as Dogfile), load values into env
	if _, err := os.Stat(`./.env`); !os.IsNotExist(err) {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	a, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if a.help {
		printHelp()
		os.Exit(0)
	}

	if a.version {
		printVersion()
		os.Exit(0)
	}

	tm, err := parser.LoadDogFile()
	if err != nil {
		printNoValidDogfile()
		os.Exit(1)
	}

	if a.taskName != "" {
		runner, err := execute.NewRunner(tm, a.info)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		runner.Run(a.taskName)
	} else {
		printTasks(tm)
		os.Exit(0)
	}
}
