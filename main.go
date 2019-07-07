package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gocmd/subTwo"
	"log"
	"os"
)

var logPath string

func cmdFun(cmd *cobra.Command, args []string) {
	fmt.Println("\n@@@test")
}

func main(){
	var cmd = &cobra.Command{
		Use:               "gocmd",
		Short:             "a testing for cmd",
		Long:              "a testing for go cmd",
		PersistentPreRunE: prepareLog,
	}

	//fmt.Println(cmd)
	var subOne = &cobra.Command{
		Use:   "subOne",
		Short: "subOne spec",
		Long:  "subOne spec long",
		Run:   cmdFun,
	}

	//args for command, supporting more than one.
	cmd.PersistentFlags().StringVarP(&logPath, "log", "l", "/dev/null", "the file where log will be written")

	cmd.AddCommand(subOne)
	cmd.AddCommand(subTwo.Cmd)
	cmd.Execute()
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