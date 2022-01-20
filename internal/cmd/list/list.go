package list

import (
	"encoding/json"
	"fmt"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmdutils"
)

type ListOptions struct {
	PerPage int
	Query   string
}

type listCmd struct {
	client client.Client
}

type ListCmd interface {
	AsJSON(opts *ListOptions) error
	AsTable(opts *ListOptions) error
}

func NewListCmd(client client.Client) ListCmd {
	return &listCmd{
		client: client,
	}
}

func (s *listCmd) AsTable(opts *ListOptions) error {
	envResponse, err := s.client.GetEnvironments()
	if err != nil {
		return err
	}

	if *envResponse.TotalCount == 0 {
		fmt.Printf("There are no environments in %s/%s\n", s.client.GetOwner(), s.client.GetRepo())
		return err
	}

	fmt.Printf(
		"Showing %d of %d environments in %s/%s\n",
		len(envResponse.Environments),
		*envResponse.TotalCount,
		s.client.GetOwner(),
		s.client.GetRepo(),
	)

	newTable(envResponse.Environments, nil)

	return nil
}

func (s *listCmd) AsJSON(opts *ListOptions) error {
	envResponse, err := s.client.GetEnvironments()
	if err != nil {
		return err
	}

	if opts.Query != "" {
		environments, err := json.Marshal(envResponse.Environments)
		if err != nil {
			return err
		}

		var data []interface{}
		err = json.Unmarshal(environments, &data)
		if err != nil {
			return err
		}

		filterResponse := cmdutils.QueryResult{}
		err = cmdutils.QueryJSON(data, &filterResponse, opts.Query)
		if err != nil {
			return fmt.Errorf("invalid query!\n%s", err)
		}

		err = cmdutils.PrettyJSON(filterResponse.Result)
		if err != nil {
			return err
		}
	} else {
		err = cmdutils.PrettyJSON(envResponse.Environments)
		if err != nil {
			return err
		}
	}

	return nil
}
