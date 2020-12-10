package service

import (
	"Week02/dao"
	"Week02/database"
	"github.com/pkg/errors"
)

type User struct {
	Id   uint
	Name string
}

var defUser = User{
	Id:   0,
	Name: "guest",
}

func GetDetail(uid uint) (User, error) {
	var u User
	data, err := dao.DetailById(uid)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return defUser, nil
		}
		return u, err
	}

	u.Id = data.ID
	u.Name = data.Name
	return u, nil
}
