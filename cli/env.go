package cli

import (
	"github.com/spf13/cobra"
	"github.com/vsoch/comp/libcomp/comp"
	"github.com/vsoch/comp/libcomp/env"
)

var envCommand = &cobra.Command{
	Use:   "env",
	Short: "Inspect the env of a container or host",
	Long: `

# inspect the environment of vanessa/salad
$ comp env vanessa/salad

# inspect the local environment
$ comp env
$ comp env .

See https://github.com/vsoch/comp/ for installation, usage, and documentation.
`,

	// Resource pak identifier
	DisableAutoGenTag: true,
	Run:               runEnv,
}

func runEnv(cmd *cobra.Command, args []string) {

	// No arguments or a present working directory . indicates local
	if len(args) == 0 || len(args) == 1 && args[0] == "." {
		environ := env.New()
		environ.Print()
	} else {
		image := args[0]
		container := comp.New(image)
		container.Env()
	}

}

func init() {
	Root.AddCommand(envCommand)
}
