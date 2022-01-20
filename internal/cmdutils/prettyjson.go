package cmdutils

import (
	"encoding/json"

	"fmt"
)

// PrettyJSON is a convenience function to pretty print a JSON string
// using json.MarshalIndent.
func PrettyJSON(o interface{}) error {
	pretty, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(pretty))

	return nil
}
