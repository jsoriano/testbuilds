package memory

type MemoryPlugin struct {
	memory map[string]string
}

func NewMemoryPlugin() *MemoryPlugin {
	p := new(MemoryPlugin)
	p.memory = make(map[string]string)
	p.memory["example"] = "foo"
	return p
}

func (p *MemoryPlugin) GetName() string {
	return "memory"
}

func (p *MemoryPlugin) Keys() []string {
	keys := make([]string, 0, len(p.memory))
	for k := range p.memory {
		keys = append(keys, k)
	}
	return keys
}

func (p *MemoryPlugin) Setting(k string) string {
	return p.memory[k]
}
