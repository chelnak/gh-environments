package cmdutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryJSON(t *testing.T) {
	queryStr := ".[0]"
	input := []interface{}{
		"foo",
		"bar",
	}

	actual := QueryResult{}
	err := QueryJSON(input, &actual, queryStr)

	want := []interface{}{"foo"}

	assert.NoError(t, err)
	assert.Equal(t, want, actual.Result)
}
