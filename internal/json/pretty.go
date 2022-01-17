package json

import (
	"encoding/json"
	"fmt"
)

func Pretty(o interface{}) {
	pretty, err := json.MarshalIndent(o, "", "  ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(pretty))
}
