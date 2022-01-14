package cmd

import (
	"fmt"
	"log"

	"github.com/chelnak/gh-environments/internal/github"
	"github.com/chelnak/gh-environments/internal/table"
	"github.com/chelnak/gh-environments/internal/tui"
	"github.com/spf13/cobra"
)

type environmentOptions struct {
	Output string
}

func RootCmd() *cobra.Command {
	opts := &environmentOptions{}

	cmd := &cobra.Command{
		Use:   "gh environments <options>",
		Short: "Manage environments for a repository",
		Long:  "Manage environments for a repository",
		RunE: func(cmd *cobra.Command, args []string) error {

			environmentResponse, err := github.GetEnvironments()
			if err != nil {
				log.Fatal(err)
				return nil
			}

			if environmentResponse.TotalCount == 0 {
				fmt.Printf("There are no environments in %s/%s\n", environmentResponse.Context.Owner, environmentResponse.Context.Repo)
				return nil
			}

			switch opts.Output {
			case "table":
				table.Print(environmentResponse)
			case "interactive":
				tui.Start(environmentResponse)
			default:
				fmt.Printf("Unknown output type: %s\n", opts.Output)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.Output, "output", "o", "table", `Output format. One of "table, interactive"`)

	return cmd
}
