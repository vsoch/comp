package cli

import (
	"github.com/spf13/cobra"
	"github.com/vsoch/comp/libcomp/comp"
)

// The submit command group
var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run a container.",
	Long: `

$ comp run vanessa/salad

See https://github.com/vsoch/comp/ for installation, usage, and documentation.
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

	container := comp.New(image)
	container.Run()

}

func init() {
	Root.AddCommand(runCommand)
}
