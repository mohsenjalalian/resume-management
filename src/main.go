package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"	
)

type Resume struct {
	Title string `json:"title"`
	Path  string `json:"path"`
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/resume_mng?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	defer db.Close()
	handleRequests()
}

func handleRequests() {
	e := echo.New()
	e.GET("/", index)
	e.POST("/", new)
	e.Logger.Fatal(e.Start(":5600"))
}

func new(c echo.Context) error {
	fl, err := c.FormFile("file")

	if err != nil {
		return err
	}

	title := c.FormValue("title")
	path := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	// Source
	src, err := fl.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// Destination
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	var resume = Resume{Title: title, Path: path}
	db.Create(&resume)

	return c.String(http.StatusOK, "success")
}

func index(c echo.Context) error {
	resume := []Resume{}
	db.Find(&resume)

	return c.JSON(http.StatusOK, resume)
}
