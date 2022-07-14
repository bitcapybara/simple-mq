package config

type Configuration struct {
	Addr string `yaml:"addr"`
}

func FromYaml(path string) *Configuration {
	return nil
}
