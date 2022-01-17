package cmd

import (
	"log"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmd/list"
	"github.com/spf13/cobra"
)

var (
	perPage      int
	outputAsJSON bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List environments for a repository",
	Long:  "List environments for a repository, optionally outputting in JSON or an interactive format.",
	Run: func(cmd *cobra.Command, args []string) {

		githubClient, err := client.NewClient()
		if err != nil {
			log.Fatal(err)
		}

		listService := list.NewListService(githubClient)
		listOpts := list.ListOptions{
			PerPage: perPage,
		}

		if outputAsJSON {
			listService.AsJSON(&listOpts)
		} else {
			// The default is always table mode
			listService.AsPaginatedTable(&listOpts)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&perPage, "limit", "l", 30, "the number of environments to show per page")
	listCmd.Flags().BoolVarP(&outputAsJSON, "json", "j", false, "Output in JSON format")
}
