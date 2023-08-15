package main

import (
	config_parser "config-parser"
	"fmt"
	wrapper "grpc-server"
	"os"
	"user-services/proto"
	"user-services/service"
)

func main() {
	config, err := config_parser.Parse("./config.yml")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	server := wrapper.NewServer(config)
	service := service.New()
	proto.RegisterUserServiceServer(server.GetRef(), service)

	// TODO: Add services here

	server.Listen()
}
