package main

import (
	"fmt"
	"gomicrotest/server/service/test/handler"

	"github.com/micro/go-micro"
)

func main() {
	fmt.Println("Hello")

	service := micro.NewService(
		micro.Name("mymicro"),
		micro.Version("last"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)
	service.Init()

	server := service.Server()
	server.Handle(server.NewHandler(&handler.TestUser{}))

	err := service.Run()
	if err != nil {
		fmt.Println(err)
	}

}
