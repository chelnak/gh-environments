package cmd

import (
	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmd"
	"github.com/spf13/cobra"
)

var (
	waitTimer            int
	reviewers            string
	protectedBranches    bool
	customBranchPolicies bool
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <environment>",
	Short: "create an environment.",
	Long:  "create an environment.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(command *cobra.Command, args []string) error {

		githubClient, err := client.NewClient()
		if err != nil {
			return err
		}

		c := cmd.NewCreateCmd(githubClient)
		opts := cmd.CreateOptions{
			Name:                 args[0],
			WaitTimer:            waitTimer,
			Reviewers:            &reviewers,
			ProtectedBranches:    &protectedBranches,
			CustomBranchPolicies: &customBranchPolicies,
		}

		err = c.CreateEnvironment(opts)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().IntVarP(&waitTimer, "wait-timer", "w", 0, "Set an amount of time (in minutes) to wait before allowing deployments to proceed.")
	createCmd.Flags().StringVarP(&reviewers, "reviewers", "r", "", "Specify people or teams that may approve workflow runs when they access this environment.")
	createCmd.Flags().BoolVarP(&protectedBranches, "protected-branches", "p", false, "Deployment limited to branches with protection rules.")
	createCmd.Flags().BoolVarP(&customBranchPolicies, "selected-branches", "s", false, "Specify a list of branches using name patterns.")
}
