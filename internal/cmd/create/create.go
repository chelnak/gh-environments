package create

import (
	"fmt"
	"strings"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/google/go-github/v42/github"
)

type CreateOptions struct {
	Name                 string
	WaitTimer            int
	Reviewers            string
	ProtectedBranches    bool
	CustomBranchPolicies bool
}

type createCmd struct {
	client client.Client
}

type CreateCmd interface {
	CreateEnvironment(opts *CreateOptions) error
}

func (s createCmd) CreateEnvironment(opts *CreateOptions) error {
	var err error
	var reviewers []*github.EnvReviewers

	if opts.Reviewers != "" {
		reviewers, err = s.processReviewers(opts.Reviewers)
		if err != nil {
			return err
		}
	}

	err = s.client.CreateEnvironment(
		opts.Name,
		opts.WaitTimer,
		reviewers,
		github.BranchPolicy{
			ProtectedBranches:    &opts.ProtectedBranches,
			CustomBranchPolicies: &opts.CustomBranchPolicies,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func NewCreateCmd(client client.Client) CreateCmd {
	return &createCmd{
		client: client,
	}
}

func (s createCmd) resolveReviewer(reviewer string) (*github.EnvReviewers, error) {
	var err error
	var entityType string
	var entityID int64
	var user *github.User
	var team *github.Team

	user, err = s.client.GetUser(reviewer)
	if err != nil {
		return nil, err
	}

	if user != nil {
		entityType = "User"
		entityID = user.GetID()
	} else {
		team, err = s.client.GetTeam(reviewer)
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

func (s createCmd) processReviewers(reviewers string) ([]*github.EnvReviewers, error) {
	var envReviewers []*github.EnvReviewers

	for _, reviewer := range strings.Split(reviewers, ",") {
		entity, err := s.resolveReviewer(reviewer)
		if err != nil {
			return nil, err
		}
		envReviewers = append(envReviewers, entity)
	}

	return envReviewers, nil
}
