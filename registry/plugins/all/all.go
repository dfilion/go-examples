// A package that imports all the checks.  This keeps us from having
// to add each package to our main.go.
//
// Note the "_" at the beginning of each import.  That means the
// package's init() will be run but not the whole package.
package all

import (
	_ "registry/plugins/pluginone"
	_ "registry/plugins/plugintwo"
)
