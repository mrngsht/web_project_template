package dep

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/prj?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnv("MYSQL_USER", "dbuser"),
		GetEnv("MYSQL_PASS", "dbpass"),
		GetEnv("MYSQL_HOST", "mysql"),
	)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = d
}
