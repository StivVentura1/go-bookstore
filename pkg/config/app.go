package config

import (
	_ "github.com/go-sql-driver/mysql" // Importa il driver MySQL
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:24091995Cudi@@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
