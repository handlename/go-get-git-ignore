package main

import (
	"fmt"
	"os"

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

	if o.Lang == "" {
		langs, err := ggi.ListLangs()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, lang := range langs {
			fmt.Println(lang)
		}
		return
	}
	err := ggi.FetchAndSave(o.Lang, o.Out)

	if err == nil {
		fmt.Printf("saved file to %s successfuly.\n", o.Out)
		os.Exit(1)
	}

	fmt.Println(err)
}
