package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// 空白引入需要添加注释
	// 空白导入应该只在主包或测试包中，或者有一个注释证明它是正确的
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(conString string) {
	db, err := gorm.Open("mysql", conString)

	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
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
