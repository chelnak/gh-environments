package list

import (
	"io"
	"os"

	"github.com/google/go-github/v42/github"
	"github.com/olekukonko/tablewriter"
)

func newTable(environments []*github.Environment, writer io.Writer) {
	if writer == nil {
		writer = io.Writer(os.Stdout)
	}

	table := tablewriter.NewWriter(writer)
	timeFormat := "2006-01-02 15:04:05"
	for _, environment := range environments {
		table.Append([]string{
			*environment.Name,
			environment.CreatedAt.Local().Format(timeFormat),
			environment.UpdatedAt.Local().Format(timeFormat),
		})
	}

	table.SetHeader([]string{"", "", ""})
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetHeaderLine(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoFormatHeaders(false)
	table.SetColumnSeparator("")
	table.SetTablePadding("  ") // two spaces
	table.SetNoWhiteSpace(true)

	table.SetColumnColor(
		nil,
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	)

	table.Render()
}
