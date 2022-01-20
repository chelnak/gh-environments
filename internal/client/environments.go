package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v42/github"
)

func (c client) GetEnvironment(name string) (*github.Environment, error) {
	ctx := context.Background()
	env, response, err := c.GitHub.Repositories.GetEnvironment(ctx, c.GetOwner(), c.GetRepo(), name)

	if response.StatusCode != http.StatusOK || err != nil {
		switch response.StatusCode {
		case http.StatusNotFound:
			return nil, fmt.Errorf("environment %s not found", name)
		default:
			return nil, fmt.Errorf("an error ocured: %s", err)
		}
	}
	return env, nil
}

func (c client) GetEnvironments() (*github.EnvResponse, error) {
	ctx := context.Background()
	envResponse, response, err := c.GitHub.Repositories.ListEnvironments(ctx, c.GetOwner(), c.GetRepo())

	if response.StatusCode != http.StatusOK || err != nil {
		return nil, fmt.Errorf("an error ocured: %s", err)
	}

	return envResponse, nil
}

func (c client) DeleteEnvironment(name string) error {
	ctx := context.Background()
	response, err := c.GitHub.Repositories.DeleteEnvironment(ctx, c.GetOwner(), c.GetRepo(), name)

	if response.StatusCode != http.StatusOK || err != nil {
		return fmt.Errorf("an error ocured: %s", err)
	}

	return nil
}
