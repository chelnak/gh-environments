package client

import (
	"github.com/chelnak/go-gh"
	"github.com/google/go-github/v42/github"
)

type Client interface {
	GetEnvironment(name string) (*github.Environment, error)
	GetEnvironments() (*github.EnvResponse, error)
	DeleteEnvironment(name string) error
	GetSecretsForEnvironment(name string, opts *github.ListOptions) (*github.Secrets, error)
	CreateEnvironment(name string, waitTimer int, reviewers []*github.EnvReviewers, deploymentBranchPolicy github.BranchPolicy) error
	GetUser(name string) (*github.User, error)
	GetTeam(name string) (*github.Team, error)
	GetOwner() string
	GetRepo() string
	GetRepoID() (int64, error)
}

type client struct {
	GitHub *github.Client
	owner  string
	repo   string
}

func (c client) GetOwner() string {
	return c.owner
}

func (c client) GetRepo() string {
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
