package executor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/xsb/dog/types"
)

var SystemExecutor *Executor

func init() {
	switch runtime.GOOS {
	case "windows":
		SystemExecutor = NewExecutor("cmd")
	default:
		SystemExecutor = NewExecutor("sh")
	}
}

func writeTempFile(dir, prefix string, data string, perm os.FileMode) (*os.File, error) {
	f, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return f, err
	}

	if err = f.Chmod(perm); err != nil {
		return f, err
	}

	_, err = f.WriteString(data)
	return f, err
}

// Executor implements standard shell executor.
type Executor struct {
	cmd string
}

// NewExecutor returns a default executor with a cmd.
func NewExecutor(cmd string) *Executor {
	return &Executor{
		cmd,
	}
}

// Exec executes the created tmp script and writes the output to the writer.
func (ex *Executor) Exec(t *types.Task, eventsChan chan types.ExecutionEvent) error {
	f, err := writeTempFile("", "dog", t.Run, 0644)
	if err != nil {
		return err
	}

	binary, err := exec.LookPath(ex.cmd)
	if err != nil {
		return err
	}

	cmd := exec.Command(binary, f.Name())

	eventsChan <- &types.TaskStartEvent{
		t.Name,
		time.Now(),
	}

	statusCode := 0
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); !ok {
				// For unknown error status codes set it to 1
				statusCode = 1
			} else {
				statusCode = waitStatus.ExitStatus()
			}
		}
		output = append(output, []byte(err.Error())...)
	}

	eventsChan <- &types.OutputEvent{
		t.Name,
		output,
	}

	eventsChan <- &types.TaskEndEvent{
		t.Name,
		time.Now(),
		statusCode,
	}

	if err := os.Remove(f.Name()); err != nil {
		return err
	}

	return nil
}
