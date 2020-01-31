package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mohsenjalalian/resume-management/database/mysql"
	"github.com/mohsenjalalian/resume-management/migration"
	"github.com/mohsenjalalian/resume-management/route"
)

func main() {
	mysql.Open()
	if mysql.Err != nil {
		log.Println("Connection Failed to Open")
	}
	migration.Migrate()
	defer mysql.Db.Close()
	route.HandleRequests()
}
