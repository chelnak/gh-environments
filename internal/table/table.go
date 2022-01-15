package table

import (
	"fmt"
	"os"

	"github.com/chelnak/gh-environments/internal/github"
	"github.com/olekukonko/tablewriter"
)

func environmentsToTable(environments []github.Environment) {
	table := tablewriter.NewWriter(os.Stdout)
	for _, environment := range environments {
		table.Append([]string{
			fmt.Sprintf("%d", environment.Id),
			environment.Name,
			environment.CreatedAt,
			environment.UpdatedAt,
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

func Print(environmentResponse github.EnvironmentResponse) {

	fmt.Printf(
		"Showing %d of %d environments in %s/%s\n\n",
		len(environmentResponse.Environments),
		environmentResponse.TotalCount,
		environmentResponse.Context.Owner,
		environmentResponse.Context.Repo,
	)

	environmentsToTable(environmentResponse.Environments)
}
