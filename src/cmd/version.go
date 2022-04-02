package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

const VERSION_FILE = ".version"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of pwd-gen app",
	Long:  `Print the current version number of pwd-gen app`,
	Run: func(cmd *cobra.Command, args []string) {
		content, err := ioutil.ReadFile(VERSION_FILE)
		if err != nil {
			fmt.Printf("error getting version. Error: %+v\n", err)
			return
		}

		fmt.Printf("pwd-gen %s", string(content))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
