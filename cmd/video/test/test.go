package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormopentracing "gorm.io/plugin/opentracing"
	"log"
	"miniTikTok/pkg/constants"
	"os"
	"time"
)

type Video struct {
	gorm.Model
	PlayUrl       string `gorm:"column:play_url;type:varchar(255);not null"`
	CoverUrl      string `gorm:"column:cover_url;type:varchar(255);not null"`
	FavoriteCount int64  `gorm:"column:favorite_count;type:bigint;not null"`
	CommentCount  int64  `gorm:"column:comment_count;type:bigint;not null"`
	Title         string `gorm:"column:title;type:varchar(255);not null"`
	AuthorId      int64  `gorm:"column:author_id;type:bigint;not null"`
}

func (u *Video) TableName() string {
	return constants.VideoTableName
}

var DB *gorm.DB

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 2 * time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info,     // Log level
			Colorful:      true,            // 彩色打印
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 newLogger,
		},
	)
	if err != nil {
		panic("failed to connect database")
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		log.Fatal(err)
	}
	//var userId int64 = 1
	res := make([]*Video, 0)
	log.Print(len(res))
	if err := DB.Find(&res).Error; err != nil {
		log.Fatal(err)
	}
	log.Println(res[0].AuthorId, "author id")
	log.Println(res[1].ID, "id")
	log.Println(res[0].Title, "title")
	log.Println(res[0].CommentCount, "comment_count")
	log.Println(res[0].FavoriteCount, "fc")
	log.Println(res[0].CoverUrl, "cu")
	log.Println(res[0].PlayUrl, "pu")
}
