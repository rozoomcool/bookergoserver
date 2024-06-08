package main

import (
	"flag"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml

	server := apiserver.NewApiServer()
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
