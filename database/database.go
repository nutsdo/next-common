package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB = NewDB()
)


func NewDB() *gorm.DB {

	db, err := gorm.Open("mysql", "root:123456@/next?charset=utf8mb4&parseTime=True&loc=Local")

	if err!=nil {
		fmt.Println("数据库连接错误:",err)
		return nil
	}

	return db
}