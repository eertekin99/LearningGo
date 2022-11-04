package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	//My local machine doesn't have mysql, so I'll be using sqlite for the sake of this project.
	//d, err := gorm.Open("mysql", "username:password/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

//"github.com/jinzhu/gorm"
//_ "github.com/jinzhu/gorm/dialects/mysql"
//_ "github.com/jinzhu/gorm/dialects/sqlite"
//_ "github.com/mattn/go-sqlite3"
