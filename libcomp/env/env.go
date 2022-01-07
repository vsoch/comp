package env

// This is an abstract backend that is shared by Podman and Docker

import (
	"os"
	"sort"
	"strings"

	"github.com/vsoch/comp/lib/logger"
)

var (
	// info prints messages in colors
	info = logger.Logger{}
)

// An enviroment holds variables (key value pairs)
type Environment struct {
	Envars map[string]string
}

func parseEnv(output string) map[string]string {
	lines := strings.Split(output, "\n")
	return parseEnvLines(lines)
}

// parseEnvLines is a shared function to parse the environment lines
func parseEnvLines(lines []string) map[string]string {

	// Create a new list of environment variables and populate it
	vars := map[string]string{}

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)

		// Don't add empty lines
		if parts[0] == "" {
			continue
		}
		if len(parts) == 1 {
			vars[parts[0]] = ""
		} else if len(parts) == 2 {
			vars[parts[0]] = parts[1]
		}
	}
	return vars
}

// getEnv gets the environment from the local machine
func getEnv() map[string]string {
	lines := os.Environ()
	return parseEnvLines(lines)
}

// Print prints the environment
func (e *Environment) Print() {

	// Ensure we print sorted
	keys := make([]string, 0, len(e.Envars))
	for key := range e.Envars {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Print into table
	table := info.Table()
	table.SetHeader([]string{"NAME    ", "VALUE"})
	for _, key := range keys {
		table.AddRow([]string{key, e.Envars[key]})
	}
	table.Print()
}

// New creates a new parsed environment on host
func New() *Environment {
	vars := getEnv()
	return &Environment{Envars: vars}
}

// Parse an existing string with an environment
func Parse(output string) *Environment {
	vars := parseEnv(output)
	return &Environment{Envars: vars}
}
