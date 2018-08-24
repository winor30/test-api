package db

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"
	"github.com/winor30/test-api/entity"
)

var con redis.Conn

func CreateClient() {
	_con, _ := redis.Dial("tcp", "redis:6379")
	con = _con
}

func Save(user *entity.User) {
	jsonStr, _ := json.Marshal(user)
	con.Do("SET", user.Name, jsonStr)
}

func Get(name string) *entity.User {
	ustr, _ := redis.String(con.Do("GET", name))
	ubyte := ([]byte)(ustr)
	u := new(entity.User)

	json.Unmarshal(ubyte, u)
	return u
}
