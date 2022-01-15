package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

func init() {
}

var rootCmd = &cobra.Command{
	Use:     "environments [command]",
	Short:   "Work with GitHub environments",
	Long:    "Work with GitHub environments",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
