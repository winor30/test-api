package db

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"
	"github.com/winor30/test-api/entity"
)

func Save(user *entity.User) {
	con, _ := redis.Dial("tcp", "redis:6379")
	defer con.Close()

	jsonStr, _ := json.Marshal(user)
	con.Do("SET", user.Name, jsonStr)
}

func Get(name string) *entity.User {
	con, _ := redis.Dial("tcp", "redis:6379")
	defer con.Close()

	ustr, _ := redis.String(con.Do("GET", name))
	ubyte := ([]byte)(ustr)
	u := new(entity.User)

	json.Unmarshal(ubyte, u)
	return u
}
