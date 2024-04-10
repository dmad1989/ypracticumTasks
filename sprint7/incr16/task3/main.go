package main

type Config struct {
	Version string
	Plugins []string
	Stat    map[string]int
}

func (cfg *Config) Clone() *Config {
	// ...
	clone := &Config{Version: cfg.Version, Stat: make(map[string]int)}

	for _, v := range cfg.Plugins {
		clone.Plugins = append(clone.Plugins, v)
	}

	for i, f := range cfg.Stat {
		clone.Stat[i] = f
	}
	return clone
}
