package main

import (
	"flag"
	"http-rest-api/internal/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to a config file")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		return err
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		return err
	}
	return nil
}
