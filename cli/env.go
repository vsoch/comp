package cli

import (
	"github.com/spf13/cobra"
	"github.com/vsoch/compenv/libcompenv/compenv"
)

var envCommand = &cobra.Command{
	Use:   "env",
	Short: "Inspect the env of a container",
	Long: `

# run an environment "pak-dev" on cluster "sherlock"
$ compenv env vanessa/salad

See https://github.com/vsoch/compenv/ for installation, usage, and documentation.
`,

	// Resource pak identifier
	Args:              cobra.MinimumNArgs(1),
	DisableAutoGenTag: true,
	Run:               runEnv,
}

func runEnv(cmd *cobra.Command, args []string) {

	// Written out just to be clear
	image := args[0]

	container := compenv.New(image)
	container.Env()

}

func init() {
	Root.AddCommand(envCommand)
}
