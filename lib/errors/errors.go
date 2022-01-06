package errors

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Client / command line errors
var (
	CommandNotFound    = errors.New("command not found")
	Unknown            = errors.New("unknown error")
	NotEnoughArguments = errors.New("not enough arguments")
	TooManyArguments   = errors.New("too many arguments")
)

// Given an error, exit on failure if not nil
func Check(err error) {
	if err != nil {
		log.Fatalf("%v \n", err)
	}
}

// Given an exit code, resolve with the correct error types
func ResolveExitCode(err error) {

	// If no error, exit with success
	if err == nil {
		os.Exit(Success)
	}

	// Print the error for the user
	fmt.Println(err)

	switch {

	// Directory or filename was not found
	case os.IsNotExist(err):
		os.Exit(PathNotFound)

	// We don't have a category for this error
	case errors.Is(err, Unknown):
		os.Exit(UnknownError)

	default:
		os.Exit(UserError)
	}
}
