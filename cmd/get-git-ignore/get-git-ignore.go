package main

import (
	"fmt"

	"github.com/handlename/go-get-git-ignore"
	"github.com/handlename/go-opts"
)

type myOpts struct {
	Lang string `flag:"lang"`
	Out  string `flag:"out" default:".gitignore"`
}

func main() {
	o := myOpts{}
	opts.Parse(&o)
	err := ggi.FetchAndSave(o.Lang, o.Out)

	if err == nil {
		fmt.Printf("saved file to %s successfuly.\n", o.Out)
	} else {
		fmt.Println(err)
	}
}
