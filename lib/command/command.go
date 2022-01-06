package command

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// Get the environment into a map
func GetEnvironment() map[string]string {
	vars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		vars[pair[0]] = pair[1]
	}
	return vars
}

// Output is a basic structure for returning output, error, and exit code
type Output struct {
	Out      string
	Err      string
	ExitCode int
}

// RunInteractive performs a syscall (e.g., for a container shell)
func RunInteractive(cmd []string, env []string) error {

	environ := os.Environ()

	// If we have environment strings, add them
	if len(env) > 0 {
		environ = append(environ, env...)
	}
	// TODO add debug print here?
	return syscall.Exec(cmd[0], cmd, environ)
}

// RunCommand runs one command and returs an error, output, and error
func RunCommand(cmd []string, env []string) (error, Output) {

	// Define the command!
	Cmd := exec.Command(cmd[0], cmd[1:]...)
	Cmd.Env = os.Environ()

	// Prepare to write to output and error streams
	var outstream, errstream bytes.Buffer
	Cmd.Stdout = &outstream
	Cmd.Stderr = &errstream

	// If we have environment strings, add them
	if len(env) > 0 {
		Cmd.Env = append(Cmd.Env, env...)
	}

	// Run the command
	err := Cmd.Run()

	// Prepare an output object to return
	output := Output{Out: outstream.String(), Err: errstream.String()}
	if err != nil {

		// Try to derive an exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			output.ExitCode = exitError.ExitCode()
		}
		return err, output
	}

	// Assume success without an error
	output.ExitCode = 0
	return nil, output
}
