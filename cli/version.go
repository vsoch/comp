package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vsoch/compenv/version"
)

// PrintVersion prints the version to the terminal
func PrintVersion() {
	fmt.Printf("pak %s\n", version.Version)
}

// init sets up the version command
func init() {
	Root.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: `Show the compenv version.`,
	Long:  `Show the compenv version.`,
	Run:   func(cmd *cobra.Command, args []string) { PrintVersion() },
}
