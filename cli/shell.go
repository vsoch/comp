package cli

import (
	"github.com/spf13/cobra"

	core "github.com/vsoch/compenv/libcompenv/compenv"
)

var (
	keepShell bool
)

// The shell command group
var shellCommand = &cobra.Command{
	Use:   "shell",
	Short: "Start an interactive session (shell) into a container.",
	Long: `

# Shell into the container to inspect manually.
$ compenv shell ubuntu

See https://github.com/vsoch/compenv/ for installation, usage, and documentation.
`,
	DisableAutoGenTag: true,

	// container
	Args: cobra.MinimumNArgs(1),

	Run: runShell,
}

// runShell is the Run set for configCommand
func runShell(cmd *cobra.Command, args []string) {
	image := args[0]	
	container := core.New(image)
	container.Shell(!keepShell)
}

func init() {
	Root.AddCommand(shellCommand)
	shellCommand.PersistentFlags().BoolVar(&keepShell, "keep", true, "don't remove container after interactive session.")
}
