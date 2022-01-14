package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "environments [command] [flags]",
		Short: "Work with GitHub environments",
		Long:  "Work with GitHub environments",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	cmd.AddCommand(listCmd())
	return cmd
}
