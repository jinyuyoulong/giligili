package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(conString string) {
	db, err := gorm.Open("mysql", conString)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	//	设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//	打开
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
	migration()
}
