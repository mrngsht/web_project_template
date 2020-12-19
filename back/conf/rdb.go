package conf

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/prj?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("MYSQL_USER", "dbuser"),
		getEnv("MYSQL_PASS", "dbpass"),
		getEnv("MYSQL_HOST", "mysql"),
	)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = d
}
