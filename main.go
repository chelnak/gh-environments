package main

import (
	"fmt"
	"os"

	"github.com/chelnak/gh-environments/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		if err != cmd.ErrSilent {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
