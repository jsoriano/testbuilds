package testbuilds

type Plugin interface {
	GetName() string
	Keys() []string
	Setting(k string) string
}

var Plugins []Plugin

func InitPlugins() {
	if len(Plugins) == 0 {
		Plugins = make([]Plugin, 0)
	}
}

func init() {
	InitPlugins()
}
