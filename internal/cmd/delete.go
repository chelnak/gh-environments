package cmd

import (
	"fmt"
	"net/http"

	"github.com/chelnak/gh-environments/internal/client"
)

type DeleteOptions struct {
	Name  string
	Force bool
}

type DeleteCmd interface {
	Delete(opts DeleteOptions) error
}

type deleteCmd struct {
	client client.Client
}

func NewDeleteCmd(client client.Client) DeleteCmd {
	return deleteCmd{
		client: client,
	}
}

func (cmd deleteCmd) Delete(opts DeleteOptions) error {
	if response, err := cmd.client.DeleteEnvironment(opts.Name); err != nil {
		switch response.StatusCode {
		case http.StatusNotFound:
			return fmt.Errorf("environment '%s' does not exist", opts.Name)
		default:
			return err
		}
	}

	return nil
}
