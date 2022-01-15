package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chelnak/gh-environments/internal/github"
	"github.com/chelnak/gh-environments/internal/json"
	"github.com/chelnak/gh-environments/internal/table"
	"github.com/chelnak/gh-environments/internal/tui"
	"github.com/spf13/cobra"
)

type listOptions struct {
	Output string
}

var opts listOptions = listOptions{}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&opts.Output, "output", "o", "table", "the output format. One of: table, interactive, json")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List environments for a repository",
	Long:  "List environments for a repository, optionally outputting in JSON or an interactive format.",
	Run: func(cmd *cobra.Command, args []string) {

		environmentResponse, err := github.GetEnvironments()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if environmentResponse.TotalCount == 0 {
			fmt.Printf("There are no environments in %s/%s\n", environmentResponse.Context.Owner, environmentResponse.Context.Repo)
		}

		switch opts.Output {
		case "table":
			table.Print(environmentResponse)
		case "interactive":
			tui.Start(environmentResponse)
		case "json":
			json.Print(environmentResponse)
		default:
			fmt.Printf("Unknown output type: %s\n", opts.Output)
			os.Exit(1)
		}
	},
}
