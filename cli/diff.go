package cli

import (
	"github.com/spf13/cobra"
	//	"github.com/vsoch/compenv/libcompenv/backend"
	//	core "github.com/vsoch/compenv/libcompenv/compenv"
)

// The diff command group
var diffCommand = &cobra.Command{
	Use:   "diff",
	Short: "Diff environments between containers.",
	Long: `

$ compenv diff <container1> <container2>

See https://github.com/vsoch/compenv/ for installation, usage, and documentation.
`,

	// Args are pieces to compare
	Args:              cobra.MinimumNArgs(2),
	DisableAutoGenTag: true,
	Run:               runDiff,
}

// runDiff is the Run set for configCommand
func runDiff(cmd *cobra.Command, args []string) {

}

func init() {
	Root.AddCommand(diffCommand)
}
