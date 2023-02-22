package dal

import (
	"miniTikTok/cmd/favorite/dal/db"
	"miniTikTok/cmd/favorite/dal/redisDb"
)

func Init() {
	db.Init()
	redisDb.Init()
}
