package main

// This example shows an application exposing the execution of Dogfile tasks
// through an HTTP endpoint.

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dogtools/dog"
)

// Dogfile object
var d dog.Dogfile

func main() {

	// Parse Dogfile from current path
	err := d.ParseFromDisk(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Launch the HTTP server
	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	// Get task name from path
	taskName := r.URL.Path[1:]

	// Generate task chain for the task named as the URL path
	var tc dog.TaskChain
	err := tc.Generate(d, taskName)
	if err != nil {
		fmt.Fprintf(w, "task chain generation failed: %s\n", err)
		os.Exit(1)
	}

	// Run task chain, HTTP client receives info about how task finished
	err = tc.Run()
	if err != nil {
		fmt.Fprintf(w, "%s failed: %s\n", taskName, err)
		os.Exit(2)
	}
	fmt.Fprintf(w, "%s finished\n", taskName)
}
