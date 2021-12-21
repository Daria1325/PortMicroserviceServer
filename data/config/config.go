package config

import "github.com/BurntSushi/toml"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	JsonPath string `toml:"json_path"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":9080",
		JsonPath: "data/ports.json",
	}
}
func NewConfigPath(configPath string) (*Config, error) {
	config := NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		return config, err
	}
	return config, nil
}
