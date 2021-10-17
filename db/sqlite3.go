package db

import (
	"log"

	"github.com/jinzhu/gorm"
)

func NewSQLite3() *gorm.DB {
	db, err := gorm.Open("sqlite3", "rss-reader.db")
	if err != nil {
		panic("failed to connect database")
	}

	err = db.Exec(schema).Error
	if err != nil {
		log.Fatal(err)
	}
	return db
}
