package cmdutils

import (
	"encoding/json"

	"fmt"
)

// PrettyJSON is a convenience function to pretty print a JSON string
// using json.MarshalIndent.
func PrettyJSON(o interface{}) {
	pretty, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(pretty))
}
