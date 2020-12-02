package dao

import (
	"database/sql"

	xerrors "github.com/pkg/errors"
)

type UserEntity struct {
	Id   int64
	Name string
}

func NewUser() *UserEntity {
	return &UserEntity{}
}

func (u *UserEntity) SelectById(id int) (*UserEntity, error) {
	if 0 >= id {
		err := sql.ErrNoRows
		return nil, xerrors.Wrapf(err, "user is not exist.")
	}
	return &UserEntity{1, "this is a name"}, nil
}
