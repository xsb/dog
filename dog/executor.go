package dog

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"syscall"
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
func (ex *Executor) Exec(t *Task, w io.Writer) error {

	if err := t.ToDisk(); err != nil {
		return err
	}

	binary, err := exec.LookPath(ex.cmd)
	if err != nil {
		return err
	}

	cmd := exec.Command(binary, t.Path)

	w.Write([]byte(" - " + t.Name + " started\n"))

	statusCode := 0
	if output, err := cmd.CombinedOutput(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); !ok {
				// For unknown error status codes set it to 1
				statusCode = 1
			} else {
				statusCode = waitStatus.ExitStatus()
			}
		}
		w.Write(output)
		w.Write([]byte("\n" + err.Error() + "\n"))
	} else {
		w.Write(output)
	}

	msg := fmt.Sprintf(" - %s finished with status code %d\n", t.Name, statusCode)
	w.Write([]byte(msg))

	if err := os.Remove(t.Path); err != nil {
		return err
	}

	return nil
}
