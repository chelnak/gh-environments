package github

import (
	"fmt"
	"log"

	"github.com/cli/go-gh"
)

type Environment struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	Login string `json:"login"`
}

func GetEnvironments() ([]Environment, error) {

	var err error

	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	currentRepository, err := gh.CurrentRepository()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	path := fmt.Sprintf("repos/%s/%s/environments", currentRepository.Owner(), currentRepository.Name())
	var environmentResponse struct {
		TotalCount   int           `json:"total_count"`
		Environments []Environment `json:"environments"`
	}
	err = client.Get(path, &environmentResponse)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return environmentResponse.Environments, nil
}
