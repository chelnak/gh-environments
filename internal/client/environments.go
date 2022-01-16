// Copyright 2021 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"context"

	"github.com/google/go-github/v42/github"
)

func (c client) GetEnvironment(name string) (*github.Environment, error) {

	ctx := context.Background()
	env, _, err := c.GitHub.Repositories.GetEnvironment(ctx, c.Owner(), c.Repo(), name)
	if err != nil {
		return nil, err
	}

	return env, nil
}

func (c client) GetEnvironments() (*github.EnvResponse, error) {

	ctx := context.Background()
	envResponse, _, err := c.GitHub.Repositories.ListEnvironments(ctx, c.Owner(), c.Repo())
	if err != nil {
		return nil, err
	}

	return envResponse, nil
}

func (c client) DeleteEnvironment(name string) error {

	ctx := context.Background()
	_, err := c.GitHub.Repositories.DeleteEnvironment(ctx, c.Owner(), c.Repo(), name)
	if err != nil {
		return err
	}

	return nil
}
