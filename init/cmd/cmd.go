package cmd

import (
	"github.com/MCprotein/crud-server/config"
)

type Cmd struct {
	config *config.Config
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	return c
}
