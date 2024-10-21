package config

type Config struct {
	OutputFile     string   `yaml:"outputFile"`
	ExcludeFiles   []string `yaml:"excludeFiles"`
	ExcludeFormats []string `yaml:"excludeFormats"`
}

func Default() *Config {
	return &Config{OutputFile: "appendix.pb.tex"}
}
