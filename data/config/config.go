package config

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
