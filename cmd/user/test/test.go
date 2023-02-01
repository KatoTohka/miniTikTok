package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	"log"
	"miniTikTok/pkg/constants"
)

type User struct {
	gorm.Model
	UserName      string `gorm:"column:username;unique;type:varchar(30);not null"`
	PassWord      string `gorm:"column:password;type:varchar(60);not null"`
	FollowCount   int64  `gorm:"column:follow_count;type:uint;not null"`   // 关注总数
	FollowerCount int64  `gorm:"column:follower_count;type:uint;not null"` // 粉丝总数
}

func (u *User) TableName() string {
	return constants.UserTableName
}

var DB *gorm.DB

func main() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic("failed to connect database")
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		log.Fatal(err)
	}
	res := make([]*User, 0)
	if err := DB.Where("id = ?", "1").Limit(1).Find(&res).Error; err != nil {
		panic("failed query")
	}

	fmt.Printf("%+v", res[0])
	//res = []*User{
	//	{
	//		UserName: "777",
	//		PassWord: "7777",
	//	},
	//}
	//DB.Create(res)
}
