package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/google/go-github/v42/github"
)

type CreateOptions struct {
	Name                 string
	WaitTimer            int
	Reviewers            *string
	ProtectedBranches    *bool
	CustomBranchPolicies *bool
}

type CreateCmd interface {
	CreateEnvironment(opts CreateOptions) error
}

type createCmd struct {
	client client.Client
}

func NewCreateCmd(client client.Client) CreateCmd {
	return createCmd{
		client: client,
	}
}

func (cmd createCmd) CreateEnvironment(opts CreateOptions) error {
	var err error
	var reviewers []*github.EnvReviewers

	verb := "created"
	if _, _, err := cmd.client.GetEnvironment(opts.Name); err == nil {
		verb = "updated"
	}

	if opts.Reviewers != nil && reviewers == nil {
		reviewers, err = cmd.processReviewers(*opts.Reviewers)
		if err != nil {
			return err
		}
	}

	_, response, err := cmd.client.CreateEnvironment(
		opts.Name,
		opts.WaitTimer,
		reviewers,
		github.BranchPolicy{
			ProtectedBranches:    opts.ProtectedBranches,
			CustomBranchPolicies: opts.CustomBranchPolicies,
		},
	)

	if err != nil {
		switch response.StatusCode {
		case http.StatusUnprocessableEntity:
			return fmt.Errorf("there was an issue with the request. Please check the documentation for more information.\n%s", err)
		default:
			return err
		}
	}

	fmt.Printf("\n%c Environment '%s' %s successfully\n\n", '\u2705', opts.Name, verb)

	return nil
}

func (cmd createCmd) resolveReviewer(reviewer string) (*github.EnvReviewers, error) {
	var err error
	var entityType string
	var entityID int64
	var user *github.User
	var team *github.Team

	user, _, err = cmd.client.GetUser(reviewer)
	if err != nil {
		return nil, err
	}

	if user != nil {
		entityType = "User"
		entityID = user.GetID()
	} else {
		team, _, err = cmd.client.GetTeam(reviewer)
		if err != nil {
			return nil, err
		}

		if team != nil {
			entityType = "Team"
			entityID = team.GetID()
		}
	}

	if user == nil && team == nil {
		return nil, fmt.Errorf("reviewer %s not found", reviewer)
	}

	return &github.EnvReviewers{
		Type: &entityType,
		ID:   &entityID,
	}, nil
}

func (cmd createCmd) processReviewers(reviewers string) ([]*github.EnvReviewers, error) {
	var envReviewers []*github.EnvReviewers

	for _, reviewer := range strings.Split(reviewers, ",") {
		entity, err := cmd.resolveReviewer(reviewer)
		if err != nil {
			return nil, err
		}
		envReviewers = append(envReviewers, entity)
	}

	return envReviewers, nil
}
