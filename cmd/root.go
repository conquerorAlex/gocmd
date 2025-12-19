package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"gocmd/subTwo"
)

var (
	logPath string
)

var rootCmd = &cobra.Command{
	Use:               "gocmd",
	Short:             "a testing for cmd",
	Long:              "a testing for go cmd",
	PersistentPreRunE: prepareLog,
}

func init() {
	// args for command, supporting more than one.
	rootCmd.PersistentFlags().StringVarP(&logPath, "log", "l", "/dev/null", "the file where log will be written")

	rootCmd.AddCommand(subOneCmd)
	rootCmd.AddCommand(subTwo.Cmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func prepareLog(cmd *cobra.Command, args []string) error {
	_, err := os.Stat(logPath)
	var f *os.File
	if os.IsNotExist(err) {
		f, err = os.Create(logPath)
	} else {
		f, err = os.OpenFile(logPath, os.O_RDWR|os.O_APPEND, 0666)
	}
	if err != nil {
		return err
	}
	log.SetOutput(f)
	return nil
}

