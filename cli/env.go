package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vsoch/comp/libcomp/comp"
	"github.com/vsoch/comp/libcomp/env"
)

var (
	envQuiet   bool
	envJson    bool
	envPretty  bool
	envOutfile string
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
	var environ *env.Environment
	if len(args) == 0 || len(args) == 1 && args[0] == "." {
		environ = env.New()
	} else {
		image := args[0]
		container := comp.New(image)
		environ = container.Env()
	}

	// Print json to terminal (quiet does not make sense here)
	if envJson {
		output := string(environ.ToJson(envPretty))
		fmt.Println(output)
	} else if !envQuiet {
		environ.Print()
	}

	if envOutfile != "" {
		environ.SaveJson(envOutfile)
	}
}

func init() {
	envCommand.Flags().BoolVarP(&envQuiet, "quiet", "q", false, "Generate output as json")
	envCommand.Flags().BoolVarP(&envJson, "json", "j", false, "Generate output as json")
	envCommand.Flags().BoolVarP(&envPretty, "pretty", "p", false, "print pretty")
	envCommand.Flags().StringVarP(&envOutfile, "outfile", "o", "", "Save output to this json file")
	Root.AddCommand(envCommand)
}
