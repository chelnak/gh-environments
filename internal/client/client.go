package client

import (
	"github.com/chelnak/go-gh"
	"github.com/google/go-github/v42/github"
)

type Client interface {
	GetEnvironment(name string) (*github.Environment, error)
	GetEnvironments() (*github.EnvResponse, error)
	DeleteEnvironment(name string) error
	Owner() string
	Repo() string
}

type client struct {
	GitHub *github.Client
	owner  string
	repo   string
}

func (c client) Owner() string {
	return c.owner
}

func (c client) Repo() string {
	return c.repo
}

func NewClient() (Client, error) {
	restClient, err := gh.RESTClient(nil)
	if err != nil {
		return nil, err
	}

	g := github.NewClient(restClient.GetHTTPClient())

	currentRepository, err := gh.CurrentRepository()
	if err != nil {
		return nil, err
	}

	client := client{GitHub: g, owner: currentRepository.Owner(), repo: currentRepository.Name()}
	return client, nil
}