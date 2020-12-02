package main

import (
	"fmt"

	"github.com/yukixia/Go-000/Week02/dao"
	"github.com/yukixia/Go-000/Week02/service"

	"github.com/pkg/errors"
)

func main() {
	//fail
	id := -1
	ser := service.NewUserService()
	user, err := ser.GetUser(id)
	if err != nil {
		fmt.Printf("origin fail:%v\n", errors.Cause(err))
		fmt.Printf("get user fail: %+v\n", err)
	} else {
		fmt.Printf("user info:%+v\n", user)
	}
	//是否 ErrNotData
	if errors.Is(err, dao.ErrNotData) {
		fmt.Println("err is not data exist")
	}

	fmt.Println("-----------------------")
	//success
	id = 100
	ser2 := service.NewUserService()
	user1, err := ser2.GetUser(id)
	if err != nil {
		fmt.Printf("origin fail:%v\n", errors.Cause(err))
		fmt.Printf("get user fail: %+v\n", err)
	} else {
		fmt.Printf("user info:%+v\n", user1)
	}

}
