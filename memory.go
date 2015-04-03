// +build memory

package testbuilds

import "github.com/jsoriano/testbuilds/memory"

func init() {
	InitPlugins()
	Plugins = append(Plugins, memory.NewMemoryPlugin())
}
