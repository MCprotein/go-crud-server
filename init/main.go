package main

import (
	"flag"

	"github.com/MCprotein/crud-server/init/cmd"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()

	cmd.NewCmd(*configPathFlag)
}
