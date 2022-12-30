package repository

import (
	"os"
)

func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&User{},
		)
	if err != nil {
		os.Exit(0)
	}
}
