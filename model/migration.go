package model

// 数据迁移
func migration() {
	//	自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
}
