package json

import (
	"encoding/json"
	"fmt"
)

func Render(o interface{}) {
	pretty, err := json.MarshalIndent(o, "", "  ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(pretty))
}
