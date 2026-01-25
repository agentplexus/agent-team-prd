package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	prdFile string
	version = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "prdtool",
	Short: "PRD Tool - CLI for managing Product Requirements Documents",
	Long: `PRD Tool is a CLI for creating, validating, scoring, and viewing
Product Requirements Documents (PRDs) using the canonical PRD JSON schema.

This tool assists AI agents and product managers in building and
interacting with PRD documents.

Examples:
  prdtool init --title "New Feature" --owner "PM Name"
  prdtool validate PRD.json
  prdtool show PRD.json
  prdtool score PRD.json
  prdtool view --type exec PRD.json
  prdtool add problem --statement "Users cannot..."
  prdtool add persona --name "Developer Dan"`,
	Version: version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&prdFile, "file", "f", "PRD.json", "PRD file path")
}

// getPRDPath returns the PRD file path from args or flag.
func getPRDPath(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return prdFile
}

// exitWithError prints an error and exits.
func exitWithError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}
