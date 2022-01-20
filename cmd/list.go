package cmd

import (
	"errors"
	"fmt"

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
	RunE: func(cmd *cobra.Command, args []string) error {

		githubClient, err := client.NewClient()
		if err != nil {
			return err
		}

		listCmd := list.NewListCmd(githubClient)
		listOpts := list.ListOptions{
			PerPage: limit,
			Query:   query,
		}

		if query != "" && !outputAsJSON {
			err := errors.New("you must specify --json to use the --query flag")
			return err
		}

		if outputAsJSON {
			listCmd.AsJSON(&listOpts)
		} else {
			fmt.Println()
			listCmd.AsTable(&listOpts)
			fmt.Println()
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&limit, "limit", "l", 30, "the number of environments to show per page")
	listCmd.Flags().BoolVarP(&outputAsJSON, "json", "j", false, "Output in JSON format")
	listCmd.Flags().StringVarP(&query, "query", "q", "", "a query string to filter environments")
}
