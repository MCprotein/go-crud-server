package config

import (
	"os"

	"github.com/naoina/toml"
)

// node .env
// env -> toml

type Config struct {
	Server struct {
		Port int64
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)
	// var c *Config

	file, errOpen := os.Open(filePath)
	errHandler(errOpen)

	errDecode := toml.NewDecoder(file).Decode(c)
	if errDecode != nil {
		panic(errDecode)
	}

	return c
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}
