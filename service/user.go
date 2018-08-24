package service

import (
	"github.com/winor30/test-api/db"
	"github.com/winor30/test-api/entity"
)

func GetUser(name string) *entity.User {
	u := db.Get(name)
	return u
}

func SaveUser(user *entity.User) {
	db.Save(user)
}
