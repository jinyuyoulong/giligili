package model

// 数据迁移
func migration() {
	//	自动迁移模式
	// docker 初始化MySQL 默认字符集为 ascii 在这里重新设置为 utf8
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Video{})

	// DB.AutoMigrate(&User{})
	// 	DB.AutoMigrate(&Video{})
}
