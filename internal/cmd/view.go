package cmd

import (
	"fmt"
	"net/http"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmdutils"
)

type ViewOptions struct {
	Name string
}

type ViewCmd interface {
	AsJSON(opts ViewOptions) error
}

type viewCmd struct {
	client client.Client
}

func NewViewCmd(client client.Client) ViewCmd {
	return viewCmd{
		client: client,
	}
}

func (cmd viewCmd) AsJSON(opts ViewOptions) error {
	envResponse, response, err := cmd.client.GetEnvironment(opts.Name)
	if response.StatusCode != http.StatusOK || err != nil {
		switch response.StatusCode {
		case http.StatusNotFound:
			return fmt.Errorf("environment '%s' does not exist", opts.Name)
		default:
			return err
		}
	}

	err = cmdutils.PrettyJSON(envResponse)
	if err != nil {
		return err
	}

	return nil
}
