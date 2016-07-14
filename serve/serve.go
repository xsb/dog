package serve

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dogtools/dog/execute"
	"github.com/dogtools/dog/parser"
	"github.com/dogtools/dog/types"
	"github.com/gorilla/pat"
)

var tm types.TaskMap

func StartServer(taskMap types.TaskMap) {
	router := pat.New()
	router.Post("/run/{taskName}", runTaskHandler)
	log.Fatal(http.ListenAndServe(":4000", router))
}

func runTaskHandler(w http.ResponseWriter, req *http.Request) {
	taskName := req.URL.Query().Get(":taskName")

	tm, err := parser.LoadDogFile()
	if err != nil {
		fmt.Println("No Valid Dogfile")
		os.Exit(1)
	}

	runner, err := execute.NewRunner(tm, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	runner.Run(taskName)

	w.WriteHeader(200)
	w.Write([]byte("Running " + taskName + "!\n"))
}
