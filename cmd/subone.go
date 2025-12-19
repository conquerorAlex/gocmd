package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func subOneRun(cmd *cobra.Command, args []string) {
	fmt.Println("\n@@@test")
}

var subOneCmd = &cobra.Command{
	Use:   "subOne",
	Short: "subOne spec",
	Long:  "subOne spec long",
	Run:   subOneRun,
}

