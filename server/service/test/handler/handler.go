package handler

import (
	"fmt"
	ptest "gomicrotest/server/proto/test"

	"golang.org/x/net/context"
)

type TestUser struct {
}

func (c *TestUser) GetUser(ctx context.Context, req *ptest.User, rsp *ptest.User) error {
	fmt.Println(req)
	fmt.Println(req.Id)
	*rsp = ptest.User{
		Id: 32131,
	}
	return nil
}
