package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vsoch/comp/libcomp/diff"
)

var (
	diffQuiet   bool
	diffJson    bool
	diffPretty  bool
	diffOutfile string
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
	if diffJson {
		output := string(differ.ToJson(diffPretty))
		fmt.Println(output)
	} else if !diffQuiet {
		differ.PrintDiff()
	}

	if diffOutfile != "" {
		differ.SaveJson(diffOutfile)
	}
}

func init() {
	diffCommand.Flags().BoolVarP(&diffQuiet, "quiet", "q", false, "Generate output as json")
	diffCommand.Flags().BoolVarP(&diffJson, "json", "j", false, "Generate output as json")
	diffCommand.Flags().BoolVarP(&diffPretty, "pretty", "p", false, "print pretty")
	diffCommand.Flags().StringVarP(&diffOutfile, "outfile", "o", "", "Save output to this json file.")
	Root.AddCommand(diffCommand)
}
