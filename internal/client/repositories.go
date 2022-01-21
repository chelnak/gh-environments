package client

import (
	"context"
	"fmt"
	"net/http"
)

func (c client) GetRepoID() (int64, error) {
	ctx := context.Background()
	repo, response, err := c.GitHub.Repositories.Get(ctx, c.GetOwner(), c.GetRepo())

	if response.StatusCode != http.StatusOK || err != nil {
		switch response.StatusCode {
		case http.StatusNotFound:
			return 0, fmt.Errorf("repository %s not found", c.GetRepo())
		default:
			return 0, fmt.Errorf("an error ocured: %s", err)
		}
	}

	return repo.GetID(), nil
}
