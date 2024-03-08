package xhandler

import (
	r "aqbiluts/response"
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDBHandler() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/utspbp?parseTime=true")
	r.CheckError(err)
	return db
}

func ConnectGormDBHandler() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/utspbp?parseTime=true"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
