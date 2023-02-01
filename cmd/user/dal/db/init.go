package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	"log"
	"miniTikTok/pkg/constants"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	if err = DB.Use(gormopentracing.New()); err != nil {
		log.Fatal(err)
	}
	//	migrate
	m := DB.Migrator()
	if m.HasTable(constants.UserTableName) {
		return
	}
	if err = m.CreateTable(&User{}); err != nil {
		log.Fatal(err)
	}
}
