package Gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func NewDB() (err error) {
	//连接数据库
	Db, err = gorm.Open("mysql", "root:13896764180zy@tcp(127.0.0.1:3306)/CLASS?"+"charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//最大空闲
	Db.DB().SetMaxIdleConns(100)
	//最大连接
	Db.DB().SetMaxOpenConns(100)
	return err
}
