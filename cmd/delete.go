/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <environment>",
	Short: "Delete an environment.",
	Long:  "Delete an environment.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var err error

		githubClient, err := client.NewClient()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		promptText := fmt.Sprintf("You are about to delete %s. Are you sure that you want to continue?", args[0])
		confirm := confirmation.New(promptText, confirmation.No)
		ready, err := confirm.RunPrompt()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if ready {
			err = githubClient.DeleteEnvironment(args[0])
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
