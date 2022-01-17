package list

import (
	"log"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/json"
)

type ListOptions struct {
	PerPage int
}

type listService struct {
	client client.Client
}

type ListService interface {
	AsJSON(opts *ListOptions)
	AsPaginatedTable(opts *ListOptions)
}

func NewListService(client client.Client) ListService {
	return &listService{
		client: client,
	}
}

func (s *listService) AsPaginatedTable(opts *ListOptions) {
	newPaginatedTable(s.client, opts.PerPage)
}

func (s *listService) AsJSON(opts *ListOptions) {
	envResponse, err := s.client.GetEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	json.Pretty(envResponse)
}
