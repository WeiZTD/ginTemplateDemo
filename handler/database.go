package handler

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetupDB() error {
	var err error
	db, err = gorm.Open("mysql", "root:password@/ptt?charset=utf8mb4&parseTime=True")
	if err != nil {
		return err
	}
	//db.LogMode(true)
	return nil
}

func CloseDB() {
	db.Close()
}
