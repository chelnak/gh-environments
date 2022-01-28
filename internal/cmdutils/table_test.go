package cmdutils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableWriter(t *testing.T) {
	tests := []struct {
		name    string
		headers []string
		rows    [][]string
		want    string
	}{
		{
			name:    "with empty headers",
			headers: []string{"", "", "", ""},
			rows: [][]string{
				{"foo", "bar", "baz", "qux"},
				{"baz", "qux", "foo", "bar"},
			},
			want: "                   \nfoo  bar  \x1b[0;32mbaz\x1b[0m  \x1b[0;90mqux\x1b[0m  \nbaz  qux  \x1b[0;32mfoo\x1b[0m  \x1b[0;90mbar\x1b[0m  \n",
		},
		{
			name:    "with populated headers",
			headers: []string{"foo", "bar", "baz", "qux"},
			rows: [][]string{
				{"foo", "bar", "baz", "qux"},
				{"baz", "qux", "foo", "bar"},
			},
			want: "foo  bar  baz  qux \nfoo  bar  \x1b[0;32mbaz\x1b[0m  \x1b[0;90mqux\x1b[0m  \nbaz  qux  \x1b[0;32mfoo\x1b[0m  \x1b[0;90mbar\x1b[0m  \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			tw := NewTableWriter(tt.headers, tt.rows, &output)
			err := tw.Write()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, output.String())
		})
	}
}
