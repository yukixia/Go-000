package service

import "github.com/yukixia/Go-000/Week02/dao"

type UserService struct {
	user *dao.UserEntity
}

func GetUser(id int) (*dao.UserEntity, error) {
	_, err := dao.SelectById(id)
	return nil, err
}
