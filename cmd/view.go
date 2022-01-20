/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmd/view"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view <environment> [flags]",
	Short: "View details about an environment.",
	Long:  "View details about an environment. Optionally output as JSON.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		githubClient, err := client.NewClient()
		if err != nil {
			log.Fatal(err)
		}

		viewCmd := view.NewViewCmd(githubClient)
		viewOpts := view.ViewOptions{
			Name: args[0],
		}

		viewCmd.AsJSON(&viewOpts)

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
}
