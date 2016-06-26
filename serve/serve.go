package serve

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dogtools/dog/execute"
	"github.com/dogtools/dog/types"

	"github.com/husobee/vestigo"
)

var tm types.TaskMap

func StartServer(taskMap types.TaskMap) {
	tm = taskMap

	router := vestigo.NewRouter()
	router.Post("/run/:taskName", runTask)

	log.Fatal(http.ListenAndServe(":4000", router))
}

func runTask(w http.ResponseWriter, r *http.Request) {
	taskName := vestigo.Param(r, "taskName")

	runner, err := execute.NewRunner(tm, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	runner.Run(taskName)

	w.WriteHeader(200)
	w.Write([]byte("Running " + taskName + "!\n"))
}
