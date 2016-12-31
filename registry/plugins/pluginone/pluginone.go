// Example check.
package pluginone

import (
	"fmt"
	"registry/plugins"
)

func pluginOne() {
	fmt.Println("pluginOne called")
}

func init() {
	// Calls plugins.Add() to register itself.
	plugins.Add("checkone", pluginOne)
}
