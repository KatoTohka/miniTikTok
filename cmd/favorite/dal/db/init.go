package db

import (
	"log"
	"miniTikTok/pkg/constants"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
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
		log.Fatal("failed to connect database", err)
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		log.Fatal(err)
	}
	//	migrate
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
	DB.AutoMigrate(&Favorite{})
}
