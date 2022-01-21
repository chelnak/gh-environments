package list

import (
	"fmt"
	"io"
	"os"

	"github.com/chelnak/gh-environments/internal/cmdutils"
	"github.com/google/go-github/v42/github"
	"github.com/olekukonko/tablewriter"
)

func (s *listCmd) newTable(environments []*github.Environment, writer io.Writer) error {
	if writer == nil {
		writer = io.Writer(os.Stdout)
	}

	table := tablewriter.NewWriter(writer)
	timeFormat := "2006-01-02 @ 15:04:05"
	for _, environment := range environments {
		secretCount, err := s.client.GetSecretsForEnvironment(*environment.Name, nil)
		if err != nil {
			return fmt.Errorf("could not get secrets for environment %s: %s", *environment.Name, err)
		}

		protectionRuleCount := len(environment.ProtectionRules)

		table.Append([]string{
			*environment.Name,
			fmt.Sprintf("%d protection %s", protectionRuleCount, cmdutils.Pluralize(protectionRuleCount, "rule", "rules")),
			fmt.Sprintf("%d %s", secretCount.TotalCount, cmdutils.Pluralize(secretCount.TotalCount, "secret", "secrets")),
			fmt.Sprintf("updated %s", environment.UpdatedAt.Local().Format(timeFormat)),
		})
	}

	table.SetHeader([]string{"", "", "", ""})
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetHeaderLine(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoFormatHeaders(false)
	table.SetColumnSeparator("")
	table.SetTablePadding("  ") // two spaces
	table.SetNoWhiteSpace(true)
	table.SetAutoWrapText(false)

	table.SetColumnColor(
		nil,
		nil,
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiBlackColor},
	)

	table.Render()

	return nil
}
