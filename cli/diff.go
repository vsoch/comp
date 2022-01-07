package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vsoch/comp/libcomp/env"
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
	var envA, envB *env.Environment
	if args[0] == "." {
		envA = env.New()
	} else if args[1] == "." {
		envB = env.New()
	}
	fmt.Println(envA)
	fmt.Println(envB)
}

func init() {
	Root.AddCommand(diffCommand)
}
