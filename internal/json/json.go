package json

import (
	"encoding/json"
	"fmt"

	"github.com/chelnak/gh-environments/internal/github"
)

func Print(environmentResponse github.EnvironmentResponse) {

	pretty, err := json.MarshalIndent(environmentResponse, "", "  ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(pretty))

}
