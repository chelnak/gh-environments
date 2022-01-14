package github

import (
	"fmt"
	"log"

	"github.com/cli/go-gh"
)

type Context struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
}

type Environment struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type EnvironmentResponse struct {
	Environments []Environment `json:"environments"`
	TotalCount   int           `json:"total_count"`
	Context      Context       `json:"context"`
}

func GetEnvironments() (EnvironmentResponse, error) {

	var err error

	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
		return EnvironmentResponse{}, err
	}

	currentRepository, err := gh.CurrentRepository()
	if err != nil {
		log.Fatal(err)
		return EnvironmentResponse{}, err
	}

	path := fmt.Sprintf("repos/%s/%s/environments", currentRepository.Owner(), currentRepository.Name())

	environmentResponse := EnvironmentResponse{}
	err = client.Get(path, &environmentResponse)
	if err != nil {
		log.Fatal(err)
		return EnvironmentResponse{}, err
	}

	environmentResponse.Context = Context{
		Owner: currentRepository.Owner(),
		Repo:  currentRepository.Name(),
	}

	return environmentResponse, nil
}
