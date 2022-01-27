package client

import (
	"context"

	"github.com/google/go-github/v42/github"
)

func (c client) GetTeam(name string) (*github.Team, *github.Response, error) {
	ctx := context.Background()
	team, response, err := c.gitHub.Teams.GetTeamBySlug(ctx, c.GetOwner(), name)

	if err != nil {
		return nil, response, err
	}

	return team, response, nil
}
