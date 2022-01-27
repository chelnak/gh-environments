package cmdutils

import (
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

type TableWriter interface {
	Write() error
}

type tableWriter struct {
	headers []string
	data    [][]string
	writer  io.Writer
}

func NewTableWriter(headers []string, data [][]string, writer io.Writer) TableWriter {
	return tableWriter{
		headers: headers,
		data:    data,
		writer:  writer,
	}
}

func (t tableWriter) Write() error {
	if t.writer == nil {
		t.writer = io.Writer(os.Stdout)
	}

	table := tablewriter.NewWriter(t.writer)

	for _, d := range t.data {
		table.Append(d)
	}

	table.SetHeader(t.headers)
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
