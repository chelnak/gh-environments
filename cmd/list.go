package cmd

import (
	"fmt"
	"log"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmd/list"
	"github.com/spf13/cobra"
)

var (
	limit        int
	outputAsJSON bool
	query        string
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

		listCmd := list.NewListCmd(githubClient)
		listOpts := list.ListOptions{
			PerPage: limit,
			Query:   query,
		}

		if query != "" && !outputAsJSON {
			fmt.Println("You must specify --json to use the --query flag")
			return
		}

		if outputAsJSON {
			listCmd.AsJSON(&listOpts)
		} else {
			listCmd.AsTable(&listOpts)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&limit, "limit", "l", 30, "the number of environments to show per page")
	listCmd.Flags().BoolVarP(&outputAsJSON, "json", "j", false, "Output in JSON format")
	listCmd.Flags().StringVarP(&query, "query", "q", "", "a query string to filter environments")
}
