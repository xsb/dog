package run

import (
	"bufio"
	"io"
	"os"
)

// Runner just runs anything.
type Runner interface {
	// StdoutPipe returns a pipe that will be connected to the runner's
	// standard output when the command starts.
	StdoutPipe() (io.ReadCloser, error)

	// StderrPipe returns a pipe that will be connected to the runner's
	// standard error when the command starts.
	StderrPipe() (io.ReadCloser, error)

	// Start starts the runner but does not wait for it to complete.
	Start() error

	// Process returns the underlying process once started
	// It can be used for forwarding signals for graceful shutdown etc.
	GetProcess() *os.Process

	// Wait waits for the runner to exit. It must have been started by Start.
	//
	// The returned error is nil if the runner has no problems copying
	// stdin, stdout, and stderr, and exits with a zero exit status.
	Wait() error
}

// NewShRunner creates a system standard shell script runner.
func NewShRunner(code string, workdir string, env []string) (Runner, error) {
	return newCmdRunner(runCmdProperties{
		runner:        "sh",
		fileExtension: ".sh",
		code:          code,
		workdir:       workdir,
		env:           env,
	})
}

// NewBashRunner creates a Bash runner.
func NewBashRunner(code string, workdir string, env []string) (Runner, error) {
	return newCmdRunner(runCmdProperties{
		runner:        "bash",
		fileExtension: ".sh",
		code:          code,
		workdir:       workdir,
		env:           env,
	})
}

// NewPythonRunner creates a Python runner.
func NewPythonRunner(code string, workdir string, env []string) (Runner, error) {
	return newCmdRunner(runCmdProperties{
		runner:        "python",
		fileExtension: ".py",
		code:          code,
		workdir:       workdir,
		env:           env,
	})
}

// NewRubyRunner creates a Ruby runner.
func NewRubyRunner(code string, workdir string, env []string) (Runner, error) {
	return newCmdRunner(runCmdProperties{
		runner:        "ruby",
		fileExtension: ".rb",
		code:          code,
		workdir:       workdir,
		env:           env,
	})
}

// NewPerlRunner creates a Perl runner.
func NewPerlRunner(code string, workdir string, env []string) (Runner, error) {
	return newCmdRunner(runCmdProperties{
		runner:        "perl",
		fileExtension: ".pl",
		code:          code,
		workdir:       workdir,
		env:           env,
	})
}

// NewNodejsRunner creates a Node.js runner.
func NewNodejsRunner(code string, workdir string, env []string) (Runner, error) {
	return newCmdRunner(runCmdProperties{
		runner:        "node",
		fileExtension: ".js",
		code:          code,
		workdir:       workdir,
		env:           env,
	})
}

// NewGoRunner creates a Go runner that uses 'go run'.
func NewGoRunner(code string, workdir string, env []string) (Runner, error) {
	return newCmdRunner(runCmdProperties{
		runner:        "go run",
		fileExtension: ".go",
		code:          code,
		workdir:       workdir,
		env:           env,
	})
}

// GetOutputs is a helper method that returns both stdout and stderr outputs
// from the runner.
func GetOutputs(r Runner) (io.Reader, io.Reader, error) {
	stdout, err := r.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	stderr, err := r.StderrPipe()
	if err != nil {
		return nil, nil, err
	}

	return bufio.NewReader(stdout), bufio.NewReader(stderr), nil
}
