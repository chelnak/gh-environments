package view

import (
	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmdutils"
)

type ViewOptions struct {
	Name string
}

type viewCmd struct {
	client client.Client
}

type ViewCmd interface {
	AsJSON(opts *ViewOptions) error
}

func (s *viewCmd) AsJSON(opts *ViewOptions) error {
	envResponse, err := s.client.GetEnvironment(opts.Name)
	if err != nil {
		return err
	}

	cmdutils.PrettyJSON(envResponse)

	return nil
}

func NewViewCmd(client client.Client) ViewCmd {
	return &viewCmd{
		client: client,
	}
}
