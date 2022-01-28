package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmdutils"
)

type ListOptions struct {
	PerPage int
	Query   string
}

type ListCmd interface {
	AsJSON(opts ListOptions) error
	AsTable(opts ListOptions) error
}

type listCmd struct {
	client client.Client
}

func NewListCmd(client client.Client) ListCmd {
	return listCmd{
		client: client,
	}
}

func (cmd listCmd) AsTable(opts ListOptions) error {
	envResponse, _, err := cmd.client.GetEnvironments()
	if err != nil {
		return err
	}

	if *envResponse.TotalCount == 0 {
		fmt.Printf("There are no environments in %s/%s\n", cmd.client.GetOwner(), cmd.client.GetRepo())
		return err
	}

	data := make([][]string, len(envResponse.Environments))

	for i, environment := range envResponse.Environments {
		secretCount, _, err := cmd.client.GetSecretsForEnvironment(*environment.Name, nil)
		// We don't want secret retrieval to fail the whole command, so we set the count to 0
		if err != nil {
			secretCount.TotalCount = 0
		}

		protectionRuleCount := len(environment.ProtectionRules)

		row := []string{
			*environment.Name,
			fmt.Sprintf("%d protection %s", protectionRuleCount, cmdutils.Pluralize(protectionRuleCount, "rule", "rules")),
			fmt.Sprintf("%d %s", secretCount.TotalCount, cmdutils.Pluralize(secretCount.TotalCount, "secret", "secrets")),
			fmt.Sprintf("updated %s", environment.UpdatedAt.Local().Format("2006-01-02 @ 15:04:05")),
		}

		data[i] = row
	}

	fmt.Printf(
		"\nShowing %d of %d environments in %s/%s\n",
		len(envResponse.Environments),
		*envResponse.TotalCount,
		cmd.client.GetOwner(),
		cmd.client.GetRepo(),
	)

	headers := []string{"", "", "", ""}
	tableWriter := cmdutils.NewTableWriter(headers, data, nil)
	if err := tableWriter.Write(); err != nil {
		return err
	}

	fmt.Println()

	return nil
}

func (cmd listCmd) AsJSON(opts ListOptions) error {
	writer := os.Stdout
	envResponse, _, err := cmd.client.GetEnvironments()
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

		err = cmdutils.PrettyJSON(writer, filterResponse.Result)
		if err != nil {
			return err
		}
	} else {
		err = cmdutils.PrettyJSON(writer, envResponse.Environments)
		if err != nil {
			return err
		}
	}

	return nil
}
