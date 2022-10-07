package models

func migration() {
	//自动迁移模式
	var models = []interface{}{&Video{}, &User{}, &Review{},
		&Interactive{}, &Comment{}, &Reply{}, &Announce{},
		&AnnounceUser{}, &Messages{}, &Danmu{}, &Carousel{},
		&Admin{}}

	Db.AutoMigrate(models...)
	// Db.Set("gorm:table_options", "charset=utf8mb4").
	// 	AutoMigrate(&User{}).
	// 	AutoMigrate(&Video{}).
	// 	AutoMigrate(&Review{}).
	// 	AutoMigrate(&Interactive{}).
	// 	AutoMigrate(&Comment{}).
	// 	AutoMigrate(&Reply{}).
	// 	AutoMigrate(&Announce{}).
	// 	AutoMigrate(&AnnounceUser{}).
	// 	AutoMigrate(&Messages{}).
	// 	AutoMigrate(&Danmu{}).
	// 	AutoMigrate(&Carousel{}).
	// 	AutoMigrate(&Admin{})
}
