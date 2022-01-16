package table

import (
	"fmt"
	"os"

	"github.com/google/go-github/v42/github"
	"github.com/olekukonko/tablewriter"
)

const timeFormat string = "2006-01-02 15:04:05"

func environmentsToTable(environments []*github.Environment) {
	table := tablewriter.NewWriter(os.Stdout)
	for _, environment := range environments {
		table.Append([]string{
			fmt.Sprintf("%d", *environment.ID),
			*environment.Name,
			environment.CreatedAt.Local().Format(timeFormat),
			environment.UpdatedAt.Local().Format(timeFormat),
		})
	}

	table.SetHeader([]string{"id", "name", "created", "updated"})

	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetHeaderLine(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoFormatHeaders(false)
	table.SetColumnSeparator("")
	table.SetTablePadding("  ") // two spaces
	table.SetNoWhiteSpace(true)

	table.SetColumnColor(nil,
		nil,
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor})

	table.Render()
}

func Render(environmentResponse github.EnvResponse) {
	environmentsToTable(environmentResponse.Environments)
}
