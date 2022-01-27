package client

import (
	"context"

	"github.com/google/go-github/v42/github"
)

func (c client) GetUser(name string) (*github.User, *github.Response, error) {
	ctx := context.Background()
	user, response, err := c.gitHub.Users.Get(ctx, name)

	if err != nil {
		return nil, response, err
	}

	return user, response, nil
}
