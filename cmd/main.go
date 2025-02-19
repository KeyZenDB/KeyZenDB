package main

import (
	"flag"

	"github.com/KeyZenDB/KeyZenDB/config"
)

func main() {
	configPath := flag.String("config", "config.toml", "config file")
	flag.Parse()
	_ = config.LoadConfig(*configPath)
}
