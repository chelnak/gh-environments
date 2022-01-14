package main

import (
	"fmt"
	"os"

	"github.com/chelnak/gh-environments/cmd"
)

func main() {

	cmd := cmd.RootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
