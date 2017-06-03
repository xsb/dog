package main

import (
	"github.com/dogtools/dog"
	"gopkg.in/urfave/cli.v2"
)

var globalFlags = [...]string{
	"info", "i", "Print execution info (duration, exit status) after task execution",
	"workdir", "w", "Specify the working directory",
	"directory", "d", "Specify the dogfiles' directory",
	"version", "v", "Print version information",
	"debug", "", "Print debug information before running tasks",
}

func parse(dogfile dog.Dogfile) *cli.App {
	app := &cli.App{}

	// Add the global dog flags
	var flags []cli.Flag
	for i := 0; i < len(globalFlags); i += 3 {
		flags = append(flags, &cli.StringFlag{
			Name:    globalFlags[i],
			Aliases: []string{globalFlags[i+1]},
			Usage:   globalFlags[i+2],
		})
	}
	app.Flags = flags

	// Add the subcommands from the dogfile
	var commands []*cli.Command
	for _, task := range dogfile.Tasks {
		var taskFlags []cli.Flag

		// Add the params for the subcommand
		for _, param := range task.Params {
			taskFlags = append(taskFlags, &cli.StringFlag{
				Name:  param.Name,
				Value: param.Default,
			})
		}

		commands = append(commands, &cli.Command{
			Name:  task.Name,
			Usage: task.Description,
			Flags: taskFlags,
		})
	}
}
