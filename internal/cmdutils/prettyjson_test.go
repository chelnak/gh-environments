package cmdutils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrettyJson(t *testing.T) {
	type Foo struct {
		Name string `json:"name"`
	}

	foo := Foo{
		Name: "bar",
	}

	expect := "{\n  \"name\": \"bar\"\n}\n"

	var output bytes.Buffer
	err := PrettyJSON(&output, foo)
	assert.NoError(t, err)
	assert.Equal(t, expect, output.String())
}
