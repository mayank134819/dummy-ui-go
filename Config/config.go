package config

type Config struct {
	ServerAddr string
}

func NewConfig() *Config {
	return &Config{
		ServerAddr: "0.0.0.0:8000",
	}
}
