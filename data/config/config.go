package config

import "github.com/BurntSushi/toml"

type Config struct {
	BindAddr   string `toml:"bind_addr"`
	JsonPath   string `toml:"json_path"`
	DbUser     string `toml:"db_user"`
	DbName     string `toml:"db_name"`
	DbPassword string `toml:"db_password"`
	DbPort     string `toml:"db_port"`
	DbHost     string `toml:"db_host"`
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
