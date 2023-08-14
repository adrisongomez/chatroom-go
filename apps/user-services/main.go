package main

import (
	config_parser "config-parser"
	"fmt"
	wrapper "grpc-server"
	"os"
)

func main() {
	config, err := config_parser.Parse("./config.yml")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	server := wrapper.NewServer(config)

	// TODO: Add services here

	server.Listen()
}
