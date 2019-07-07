package subTwo

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "subTwo",
	Short: "subTwo spec",
	Long:  "testing for subTwo",
}

func init() {
	Cmd.AddCommand(initCmd)
}
