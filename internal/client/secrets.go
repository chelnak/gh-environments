package client

import (
	"context"
	"fmt"

	"github.com/google/go-github/v42/github"
)

func (c client) GetSecretsForEnvironment(environmentName string, opts *github.ListOptions) (*github.Secrets, error) {
	ctx := context.Background()
	repoID, err := c.GetRepoID()
	if err != nil {
		return nil, fmt.Errorf("could not get repository id for %s/%s\n%s", c.GetOwner(), c.GetRepo(), err)
	}

	secrets, _, err := c.GitHub.Actions.ListEnvSecrets(ctx, int(repoID), environmentName, opts)
	if err != nil {
		return nil, fmt.Errorf("could not get repository secrets for environment %s: %s", environmentName, err)
	}

	return secrets, nil
}
