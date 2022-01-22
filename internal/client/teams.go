package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v42/github"
)

func (c client) GetTeam(name string) (*github.Team, error) {
	ctx := context.Background()
	team, response, err := c.GitHub.Teams.GetTeamBySlug(ctx, c.GetOwner(), name)

	if response.StatusCode != http.StatusOK || err != nil {
		switch response.StatusCode {
		case http.StatusNotFound:
			return nil, nil
		default:
			return nil, fmt.Errorf("an error ocured: %s", err)
		}
	}
	return team, nil
}
