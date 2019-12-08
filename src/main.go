package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

type Resume struct {
	Id    int
	Title string
	Path  string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/resume_mng?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	  }
	defer db.Close()
	handleRequests()
}

func handleRequests() {
	e := echo.New()
	e.POST("/", index)
	e.Logger.Fatal(e.Start(":5500"))
}

func index(c echo.Context) error {
	fl, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// Source
	src, err := fl.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// Destination
	dst, err := os.Create(fl.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "success")
}
