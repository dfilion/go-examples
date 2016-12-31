package main

import (
	"fmt"
	"registry/plugins"

	// Import but only initialize the package.  Using "_" prevents the
	// "imported but not used" error from happening.
	_ "registry/plugins/all"
)

func main() {

	fmt.Printf("%d\n", len(plugins.Plugins))

	if len(plugins.Plugins) > 0 {
		for k,v := range plugins.Plugins {
			fmt.Printf("Found %s\n", k)
			v()
		}
	} else {
		fmt.Println("Empty")
	}
}
