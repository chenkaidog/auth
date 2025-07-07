package db

import (
	"auth/biz/db/mysql"
	"auth/biz/db/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
