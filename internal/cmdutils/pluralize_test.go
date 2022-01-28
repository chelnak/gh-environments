package cmdutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluralize(t *testing.T) {
	singular := "foo"
	plural := "foos"

	test := []struct {
		name  string
		count int
		want  string
	}{
		{
			name:  "count=0",
			count: 0,
			want:  "foos",
		},
		{
			name:  "count=1",
			count: 1,
			want:  "foo",
		},
		{
			name:  "count=2",
			count: 2,
			want:  "foos",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			actual := Pluralize(tt.count, singular, plural)
			assert.Equal(t, tt.want, actual)
		})
	}
}
