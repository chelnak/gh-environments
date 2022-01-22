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

func (c client) CreateEnvironment(name string, waitTimer int, reviewers []*github.EnvReviewers, deploymentBranchPolicy github.BranchPolicy) error {
	ctx := context.Background()

	opts := github.CreateUpdateEnvironment{
		WaitTimer: &waitTimer,
		Reviewers: reviewers,
	}

	if *deploymentBranchPolicy.CustomBranchPolicies || *deploymentBranchPolicy.ProtectedBranches {
		opts.DeploymentBranchPolicy = &deploymentBranchPolicy
	}

	_, response, err := c.GitHub.Repositories.CreateUpdateEnvironment(ctx, c.GetOwner(), c.GetRepo(), name, &opts)

	if response.StatusCode != http.StatusOK || err != nil {
		switch response.StatusCode {
		case http.StatusUnprocessableEntity:
			return fmt.Errorf("there was an issue with the request. Please check the documentation for more information.\n%s", err)
		default:
			return fmt.Errorf("an error ocured: %s", err)
		}
	}

	return nil
}

func (c client) DeleteEnvironment(name string) error {
	ctx := context.Background()
	_, err := c.GitHub.Repositories.DeleteEnvironment(ctx, c.GetOwner(), c.GetRepo(), name)

	if err != nil {
		return fmt.Errorf("an error ocured: %s", err)
	}

	return nil
}
