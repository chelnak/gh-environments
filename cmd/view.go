package cmd

import (
	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmd"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view <environment> [flags]",
	Short: "View details about an environment.",
	Long:  "View details about an environment. Optionally output as JSON.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(command *cobra.Command, args []string) error {

		githubClient, err := client.NewClient()
		if err != nil {
			return err
		}

		viewCmd := cmd.NewViewCmd(githubClient)
		viewOpts := cmd.ViewOptions{
			Name: args[0],
		}

		err = viewCmd.AsJSON(viewOpts)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
