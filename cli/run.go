package cli

import (
	"github.com/spf13/cobra"
	"github.com/vsoch/compenv/libcompenv/compenv"
)

// The submit command group
var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run a pak on a cluster.",
	Long: `

# run an environment "pak-dev" on cluster "sherlock"
$ pak run --cpus 6 --memory 2GB sherlock:pak-dev

See https://github.com/vsoch/compenv/ for installation, usage, and documentation.
`,

	// Resource pak identifier
	Args:              cobra.MinimumNArgs(1),
	DisableAutoGenTag: true,
	Run:               runRun,
}

// runRun is the Run set for runCommand
func runRun(cmd *cobra.Command, args []string) {

	// Written out just to be clear
	image := args[0]

	container := compenv.New(image)
	container.Run()

}

func init() {
	Root.AddCommand(runCommand)
}
