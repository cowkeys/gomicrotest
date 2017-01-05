package main

import (
	"fmt"

	ptest "gomicrotest/server/proto/test"

	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

func runClient(service micro.Service) {
	testmymicro := ptest.NewTestUserClient("mymicro", service.Client())
	rsp, err := testmymicro.GetUser(context.TODO(), &ptest.User{Id: 123, UserName: "test232"})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print response
	fmt.Println("username", rsp.Id)
}
func main() {
	service := micro.NewService()
	service.Init()

	runClient(service)
	//ptest.RegisterTestUserHandler(service.Server(), new(ptest.TestUser))

	// Run the server
	// if err := service.Run(); err != nil {
	// 	fmt.Println(err)
	// }
}
