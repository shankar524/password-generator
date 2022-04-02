package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pwd-gen",
	Short: "pwd-gen is cli tool to generate random text",
	Long: `pwd-gen is command line application that allows to generate random texts.
         Texts can be generated based on input rules`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
