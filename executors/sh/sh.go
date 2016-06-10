// Package sh provide an implementation
// for sh command executor.
package sh

import (
	"bufio"
	"io"
	"os"
	"os/exec"

	"github.com/xsb/dog/dog"
)

// Sh imlements sh shell executor.
type Sh struct {
	cmd string
}

func init() {
	sh := &Sh{}
	sh.cmd = "sh"
	dog.RegisterExecutor("sh", sh)
}

// Exec executes the created tmp script and writes the output to the writer.
func (sh *Sh) Exec(t *dog.Task, w io.Writer) error {
	if err := t.ToDisk(); err != nil {
		panic(err)
	}

	defer func() {
		// Remove temporary script
		err := os.Remove(t.Path)
		if err != nil {
			panic(err)
		}
	}()

	binary, err := exec.LookPath(sh.cmd)
	if err != nil {
		return err
	}

	cmd := exec.Command(binary, string(t.Run))
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// Collect and print STDOUT
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			w.Write(scanner.Bytes())
		}
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
