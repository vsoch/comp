package cli

import (
	"github.com/vsoch/compenv/lib/errors"

	// import all backends for containers
	_ "github.com/vsoch/compenv/backends"
)

// Main is the entrypoint to running the client
func Main() {

	// Run the root command, show any errors!
	if err := Root.Execute(); err != nil {
		errors.ResolveExitCode(err)
	}
}
