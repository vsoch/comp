package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vsoch/comp/lib/errors"
	"os"
)

var (
	showVersion bool
	configFile  string
)

// The root command group
var Root = &cobra.Command{
	Use:   "comp",
	Short: "Show help for comp commands.",
	Long: `
Comp is a simple tool to inspect and compare container 📦️ environments.
`,
	DisableAutoGenTag: true,
	Run:               runRoot,
}

// runRoot is the Run command set for Root
func runRoot(cmd *cobra.Command, args []string) {
	if showVersion {
		PrintVersion()
		errors.ResolveExitCode(nil)
	} else {
		_ = cmd.Usage()
		if len(args) > 0 {
			_, _ = fmt.Fprintf(os.Stderr, "Command not found.\n")
		}
		errors.ResolveExitCode(errors.CommandNotFound)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// A a version flag to root
	Root.Flags().BoolVarP(&showVersion, "version", "V", false, "Print the version of pak")

	// Path to a custom config file
	Root.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.pak/XXXX.conf)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println("TODO set config stuffs here")
}
