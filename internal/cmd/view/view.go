package view

import (
	"log"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/json"
)

type ViewOptions struct {
	Name string
}

type viewService struct {
	client client.Client
}

type ViewService interface {
	AsJSON(opts *ViewOptions)
}

func (s *viewService) AsJSON(opts *ViewOptions) {
	envResponse, err := s.client.GetEnvironment(opts.Name)
	if err != nil {
		log.Fatal(err)
	}

	json.Pretty(envResponse)
}

func NewViewService(client client.Client) ViewService {
	return &viewService{
		client: client,
	}
}
