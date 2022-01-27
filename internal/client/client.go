package client

import (
	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
)

type Client interface {
	GetEnvironment(name string) (*github.Environment, *github.Response, error)
	GetEnvironments() (*github.EnvResponse, *github.Response, error)
	DeleteEnvironment(name string) (*github.Response, error)
	GetSecretsForEnvironment(name string, opts *github.ListOptions) (*github.Secrets, *github.Response, error)
	CreateEnvironment(name string, waitTimer int, reviewers []*github.EnvReviewers, deploymentBranchPolicy github.BranchPolicy) (*github.Environment, *github.Response, error)
	GetUser(name string) (*github.User, *github.Response, error)
	GetTeam(name string) (*github.Team, *github.Response, error)
	GetOwner() string
	GetRepo() string
	GetRepoID() (int64, error)
}

type client struct {
	gitHub *github.Client
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
	httpClient, err := gh.HTTPClient(nil)
	if err != nil {
		return nil, err
	}

	g := github.NewClient(httpClient)

	currentRepository, err := gh.CurrentRepository()
	if err != nil {
		return nil, err
	}
	client := client{gitHub: g, owner: currentRepository.Owner(), repo: currentRepository.Name()}
	return client, nil
}
