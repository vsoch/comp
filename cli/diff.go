package cli

import (
	"github.com/spf13/cobra"
	"github.com/vsoch/comp/libcomp/diff"
)

// The diff command group
var diffCommand = &cobra.Command{
	Use:   "diff",
	Short: "Diff environments between containers.",
	Long: `

$ comp diff <container1> <container2>

See https://github.com/vsoch/comp/ for installation, usage, and documentation.
`,

	// Args are pieces to compare
	Args:              cobra.MinimumNArgs(2),
	DisableAutoGenTag: true,
	Run:               runDiff,
}

// runDiff is the Run set for configCommand
func runDiff(cmd *cobra.Command, args []string) {

	// No arguments or a present working directory . indicates local
	differ := diff.NewDiffer(args[0], args[1])
	differ.PrintDiff()
}

func init() {
	Root.AddCommand(diffCommand)
}
