package view

import (
	"log"

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
	AsJSON(opts *ViewOptions)
}

func (s *viewCmd) AsJSON(opts *ViewOptions) {
	envResponse, err := s.client.GetEnvironment(opts.Name)
	if err != nil {
		log.Fatal(err)
	}

	cmdutils.PrettyJSON(envResponse)
}

func NewViewCmd(client client.Client) ViewCmd {
	return &viewCmd{
		client: client,
	}
}
