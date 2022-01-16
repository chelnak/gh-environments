/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/json"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view <environment> [flags]",
	Short: "View details about an environment.",
	Long:  "View details about an environment. Optionally output as JSON.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var err error

		githubClient, err := client.NewClient()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		environment, err := githubClient.GetEnvironment(args[0])
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if cmd.Flag("json").Changed {

			json.Render(*environment)
		} else {
			fmt.Println("Use --json for now..")
		}

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
}
