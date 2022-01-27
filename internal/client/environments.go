package client

import (
	"context"

	"github.com/google/go-github/v42/github"
)

func (c client) GetEnvironment(name string) (*github.Environment, *github.Response, error) {
	ctx := context.Background()
	env, response, err := c.gitHub.Repositories.GetEnvironment(ctx, c.GetOwner(), c.GetRepo(), name)

	if err != nil {
		return nil, response, err
	}

	return env, response, nil
}

func (c client) GetEnvironments() (*github.EnvResponse, *github.Response, error) {
	ctx := context.Background()
	envResponse, response, err := c.gitHub.Repositories.ListEnvironments(ctx, c.GetOwner(), c.GetRepo())

	if err != nil {
		return nil, response, err
	}

	return envResponse, response, nil
}

func (c client) CreateEnvironment(name string, waitTimer int, reviewers []*github.EnvReviewers, deploymentBranchPolicy github.BranchPolicy) (*github.Environment, *github.Response, error) {
	ctx := context.Background()

	opts := github.CreateUpdateEnvironment{
		WaitTimer: &waitTimer,
		Reviewers: reviewers,
	}

	if *deploymentBranchPolicy.CustomBranchPolicies || *deploymentBranchPolicy.ProtectedBranches {
		opts.DeploymentBranchPolicy = &deploymentBranchPolicy
	}

	env, response, err := c.gitHub.Repositories.CreateUpdateEnvironment(ctx, c.GetOwner(), c.GetRepo(), name, &opts)

	if err != nil {
		return nil, response, err
	}

	return env, response, nil
}

func (c client) DeleteEnvironment(name string) (*github.Response, error) {
	ctx := context.Background()
	response, err := c.gitHub.Repositories.DeleteEnvironment(ctx, c.GetOwner(), c.GetRepo(), name)

	if err != nil {
		return response, err
	}

	return response, err
}
