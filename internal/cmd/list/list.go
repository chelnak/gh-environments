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
	AsJSON(opts *ListOptions)
	AsTable(opts *ListOptions)
}

func NewListCmd(client client.Client) ListCmd {
	return &listCmd{
		client: client,
	}
}

func (s *listCmd) AsTable(opts *ListOptions) {
	envResponse, err := s.client.GetEnvironments()
	if err != nil {
		fmt.Println(err)
		return
	}

	if *envResponse.TotalCount == 0 {
		fmt.Printf("There are no environments in %s/%s\n", s.client.GetOwner(), s.client.GetRepo())
		return
	}

	fmt.Printf(
		"Showing %d of %d environments in %s/%s\n",
		len(envResponse.Environments),
		*envResponse.TotalCount,
		s.client.GetOwner(),
		s.client.GetRepo(),
	)

	newTable(envResponse.Environments, nil)
}

func (s *listCmd) AsJSON(opts *ListOptions) {
	envResponse, err := s.client.GetEnvironments()
	if err != nil {
		fmt.Println(err)
		return
	}

	if opts.Query != "" {
		environments, err := json.Marshal(envResponse.Environments)
		if err != nil {
			fmt.Println(err)
			return
		}

		var data []interface{}
		err = json.Unmarshal(environments, &data)
		if err != nil {
			fmt.Println(err)
			return
		}
		filterResponse := cmdutils.QueryResult{}
		err = cmdutils.QueryJSON(data, &filterResponse, opts.Query)
		if err != nil {
			fmt.Println("Invalid query!\n", err)
		}

		cmdutils.PrettyJSON(filterResponse.Result)
	} else {
		cmdutils.PrettyJSON(envResponse.Environments)
	}
}
