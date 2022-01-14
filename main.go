package main

import (
	"fmt"
	"os"

	"github.com/chelnak/gh-environments/cmd"
)

func main() {

	cli := cmd.RootCmd()
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
