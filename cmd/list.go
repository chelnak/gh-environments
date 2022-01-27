package cmd

import (
	"errors"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmd"
	"github.com/spf13/cobra"
)

var (
	limit        int
	outputAsJSON bool
	query        string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List environments for a repository",
	Long:    "List environments for a repository, optionally outputting in JSON or an interactive format.",
	Aliases: []string{"ls"},
	RunE: func(command *cobra.Command, args []string) error {

		githubClient, err := client.NewClient()
		if err != nil {
			return err
		}

		listCmd := cmd.NewListCmd(githubClient)
		listOpts := cmd.ListOptions{
			PerPage: limit,
			Query:   query,
		}

		if query != "" && !outputAsJSON {
			err := errors.New("you must specify --json to use the --query flag")
			return err
		}

		if outputAsJSON {
			err = listCmd.AsJSON(listOpts)
			if err != nil {
				return err
			}
		} else {
			err = listCmd.AsTable(listOpts)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&limit, "limit", "l", 30, "the number of environments to show per page")
	listCmd.Flags().BoolVarP(&outputAsJSON, "json", "j", false, "Output in JSON format")
	listCmd.Flags().StringVarP(&query, "query", "q", "", "a query string to filter environments")
}
