package redisDb

import (
	"context"
	"fmt"
)

// 一个video一个set

// RDBSetFavorite 点赞功能开发
// videoId 需要点赞的ID   userid 用户ID
func RDBSetFavorite(videoId int64, userid int64) (bool, error) {
	var keys string = getKey(videoId)
	res, err := RDB.SIsMember(context.Background(), keys, userid).Result()
	if err != nil {
		return false, err
	}
	if res == true {
		return true, nil
	}
	_, err = RDB.SAdd(context.Background(), keys, userid).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

// RDBCancelFavorite 取消点赞
// videoId 需要点赞的ID   userid 用户ID
func RDBCancelFavorite(videoId int64, userid int64) (bool, error) {
	var keys string = getKey(videoId)
	res, err := RDB.SIsMember(context.Background(), keys, userid).Result()
	if err != nil {
		return false, err
	}
	if res == false {
		return true, nil
	}
	_, err = RDB.SRem(context.Background(), keys, userid).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

// RDBQueryFavoriteById 查询是否已经点赞了
func RDBQueryFavoriteById(videoId int64, userid int64) (bool, error) {
	var keys string = getKey(videoId)
	n, err := RDB.Exists(context.Background(), keys).Result()
	if err != nil {
		return false, err
	}
	if n == 0 {
		return false, nil
	}
	res, err := RDB.SIsMember(context.Background(), keys, userid).Result()
	if err != nil {
		return false, err
	}
	return res, nil
}

// RDBFavoriteCount 点赞数量
func RDBFavoriteCount(videoId int64) (int64, error) {
	var keys string = getKey(videoId)
	count, err := RDB.SCard(context.Background(), keys).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func getKey(videoId int64) string {
	return fmt.Sprintf("%d", videoId)
}
