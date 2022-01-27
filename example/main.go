package main

import (
	"fmt"

	"github.com/cli/go-gh"
)

func main() {
	a, err := gh.HTTPClient(nil)
	if err != nil {
		return
	}

	r, err := a.Get("https://api.github.com/repos/cli/go-gh/releases/latest")
	if err != nil {
		return
	}
	fmt.Println(r.Request.URL)
	fmt.Println(r.Request.Header)
}
