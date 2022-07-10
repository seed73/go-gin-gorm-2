package models

import (
	"fmt"

	// go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SetupDB : initializing mysql database
func SetupDBDefault() *gorm.DB {
	USER := ""
	PASS := ""
	HOST := ""
	PORT := ""
	DBNAME := ""
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
