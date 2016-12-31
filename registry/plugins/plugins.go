package plugins

import (
	"fmt"
)

// Dictionary that acts as the registry.
var Plugins = map[string]func(){}

// Public function for adding to the registry.
func Add(name string, fn func()) {
	fmt.Printf("Adding %s\n", name)
	Plugins[name] = fn
}
