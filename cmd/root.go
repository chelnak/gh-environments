package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var version = "dev"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "environments [command]",
	Aliases: []string{"env"},
	Short:   "Work with GitHub environments",
	Long:    "Work with GitHub environments",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
