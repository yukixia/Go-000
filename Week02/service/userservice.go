package service

import "github.com/yukixia/Go-000/Week02/dao"

// type UserService struct {
// 	user *dao.UserEntity
// }

func GetUser(id int) (*dao.UserEntity, error) {
	user := &dao.UserEntity{}
	_, err := user.SelectById(id)
	return nil, err
}
