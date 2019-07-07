package subTwo

import (
	"fmt"
	"github.com/spf13/cobra"
)
var (
	mode   string
	tmp  string
)

var initCmd = &cobra.Command{
	Use:               "init",
	Short:             "initialize a testing platform",
	Long:              "initialize a testing platform",
	PersistentPreRunE: initNs,
}

func initNs(cmd *cobra.Command, args []string) error {
	fmt.Printf("initNs\n")
	return nil
}

func initCA(cmd *cobra.Command, args []string) {
	fmt.Println("initCA")
}

var initCACmd = &cobra.Command{
	Use:   "ca",
	Short: "initialize ca",
	Long:  "initialize ca",
	Run:   initCA,
}

func init() {
	//u can add two values for one cmd
	initCACmd.Flags().StringVarP(&mode, "arg1", "m", "local","arg1")
	initCACmd.Flags().StringVarP(&tmp, "arg2", "t","./template", "arg2")

	initCmd.AddCommand(initCACmd)

	//also can add another cmd
	//initCmd.AddCommand(initMon)
}

