package config

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	ServerInfo struct {
		Port string `json:"port"`
		Info string `json:"info"`
	}
}

func NewConfig(path string) *Config {
	c := new(Config)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err := toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
