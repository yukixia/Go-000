package service

import "github.com/yukixia/Go-000/Week02/dao"

type UserService struct {
	dao *dao.UserEntity
}

func NewUserService() *UserService {
	return &UserService{dao: dao.NewUser()}
}

func (this *UserService) GetUser(id int) (*dao.UserEntity, error) {
	return nil, this.dao.SelectById(id)
}
