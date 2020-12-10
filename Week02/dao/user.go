package dao

import (
	"Week02/database"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type User struct {
	gorm.Model
	Name string
}

func (User) TableName() string {
	return `users`
}

func DetailById(uid uint) (User, error) {
	var u User
	u.ID = uid

	if err := database.GetConn().First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u, database.ErrRecordNotFound
		}
		return u, errors.Wrap(err, `query user info failed`)
	}
	return u, nil
}
