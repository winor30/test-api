package service

import (
	"github.com/winor30/test-api/entity"
)

func GetUser(name string) *entity.User {
	profile := name
	u := &entity.User{
		Name:    name,
		Profile: profile,
	}
	return u
}

func SaveUser(user *entity.User) {

}
