package dal

import (
	"log"
	"miniTikTok/cmd/feed/dal/db"
	"miniTikTok/cmd/feed/dal/redisDb"
)

func Init() {
	db.Init()
	redisDb.Init()
	log.Println("redisDb.RDB == nil : ", redisDb.RDB)
}
