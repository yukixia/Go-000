package dao

import (
	"fmt"
	"testing"
)

func Testuser(t *testing.T) {
	id := -1
	userdata := &UserEntity{}
	_, err := userdata.SelectById(id)
	fmt.Printf("%+v", err)
}
