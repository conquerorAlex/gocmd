package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// These can be set via -ldflags:
// -X 'gocmd/cmd.Version=v1.2.3' -X 'gocmd/cmd.Commit=abc123' -X 'gocmd/cmd.Date=2025-12-19'
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gocmd %s (commit=%s date=%s) %s/%s go=%s\n",
			Version, Commit, Date, runtime.GOOS, runtime.GOARCH, runtime.Version(),
		)
	},
}

