package main

import (
	"fmt"
	"log"

	"github.com/chelnak/gh-environments/github"
)

func main() {

	environments, err := github.GetEnvironments()
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(environments) == 0 {
		fmt.Println("It looks like there are no environments in this repository yet!")
		return
	}

	fmt.Println(environments)

}
