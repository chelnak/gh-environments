package cmdutils

import (
	"encoding/json"
	"io"

	"fmt"
)

// PrettyJSON is a convenience function to pretty print a JSON string
// using json.MarshalIndent.
func PrettyJSON(w io.Writer, o interface{}) error {
	pretty, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err
	}

	fmt.Fprintln(w, string(pretty))

	return nil
}
