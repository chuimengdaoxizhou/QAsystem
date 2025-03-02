package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// 获取配置文件
	con := GetConfig()
	// 获取连接数据库的dsn
	dsn := con.GetDatabasedsn()
	// 连接mysql数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	fmt.Println("Database connection established.")
	return db
}
