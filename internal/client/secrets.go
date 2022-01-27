package client

import (
	"context"

	"github.com/google/go-github/v42/github"
)

func (c client) GetSecretsForEnvironment(environmentName string, opts *github.ListOptions) (*github.Secrets, *github.Response, error) {
	ctx := context.Background()
	repoID, err := c.GetRepoID()
	if err != nil {
		return nil, nil, err
	}

	secrets, response, err := c.gitHub.Actions.ListEnvSecrets(ctx, int(repoID), environmentName, opts)

	if err != nil {
		return nil, response, err
	}

	return secrets, response, nil
}
