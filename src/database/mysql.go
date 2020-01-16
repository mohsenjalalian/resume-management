package mysql

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var Db *gorm.DB
var Err error

func Open() {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error reading .env file")
	}

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DBNAME"],
	)

	Db, Err = gorm.Open("mysql", mysqlCredentials)
}
