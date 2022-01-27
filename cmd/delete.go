package cmd

import (
	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmd"
	"github.com/spf13/cobra"
)

var force bool

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <environment>",
	Short: "Delete an environment.",
	Long:  "Delete an environment.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(command *cobra.Command, args []string) error {

		githubClient, err := client.NewClient()
		if err != nil {
			return err
		}

		deletCmd := cmd.NewDeleteCmd(githubClient)
		deleteOpts := cmd.DeleteOptions{
			Name:  args[0],
			Force: force,
		}

		err = deletCmd.Delete(deleteOpts)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&force, "force", "f", false, "Does not prompt for confirmation upon deletion")
}
