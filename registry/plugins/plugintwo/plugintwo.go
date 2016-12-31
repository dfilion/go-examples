// Example check.
package plugintwo

import (
	"fmt"
	"registry/plugins"
)

func pluginTwo() {
	fmt.Println("pluginTwo called")
}

func init() {
	// Calls plugins.Add() to register itself.
	plugins.Add("plugintwo", pluginTwo)
}
