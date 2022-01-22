package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v42/github"
)

func (c client) GetUser(name string) (*github.User, error) {
	ctx := context.Background()
	user, response, err := c.GitHub.Users.Get(ctx, name)

	if response.StatusCode != http.StatusOK || err != nil {
		switch response.StatusCode {
		case http.StatusNotFound:
			return nil, nil
		default:
			return nil, fmt.Errorf("an error ocured: %s", err)
		}
	}
	return user, nil
}
