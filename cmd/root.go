package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var version = "dev"
var ErrSilent = errors.New("ErrSilent")

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "environments [command]",
	Aliases:       []string{"env"},
	Short:         "Work with GitHub environments",
	Long:          "Work with GitHub environments",
	Version:       version,
	SilenceErrors: true,
	SilenceUsage:  true,
	Run:           nil,
}

func init() {
	RootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.Println(err)
		cmd.Println(cmd.UsageString())
		return ErrSilent
	})
}

// func Execute() error {
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
