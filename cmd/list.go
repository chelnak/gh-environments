package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/json"
	"github.com/chelnak/gh-environments/internal/table"
	"github.com/chelnak/gh-environments/internal/tui"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List environments for a repository",
	Long:  "List environments for a repository, optionally outputting in JSON or an interactive format.",
	Run: func(cmd *cobra.Command, args []string) {

		var err error

		githubClient, err := client.NewClient()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		environmentResponse, err := githubClient.GetEnvironments()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if *environmentResponse.TotalCount == 0 {
			fmt.Printf("There are no environments in %s/%s\n", githubClient.Owner(), githubClient.Repo())
			return
		}

		if cmd.Flag("json").Changed && cmd.Flag("interactive").Changed {
			fmt.Println("You can only use one of --json or --interactive")
			os.Exit(1)
		}

		if cmd.Flag("json").Changed {
			json.Render(*environmentResponse)
		} else if cmd.Flag("interactive").Changed {
			tui.Render(*environmentResponse)
		} else {
			// The default is always table mode

			fmt.Printf(
				"Showing %d of %d environments in %s/%s\n\n",
				len(environmentResponse.Environments),
				*environmentResponse.TotalCount,
				githubClient.Owner(),
				githubClient.Repo(),
			)

			table.Render(*environmentResponse)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
	listCmd.Flags().BoolP("interactive", "i", false, "Interactive mode")
}
