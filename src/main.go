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
	"code.sajari.com/docconv"
	"github.com/mohsenjalalian/resume-management/db"
)

var db *gorm.DB
var err error

type Resume struct {
	Title   string `json:"title"`
	Path    string `json:"path"`
	Content string `json:"content"`
}

func main() {
	db, err = mysql.Open()
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	defer db.Close()
	handleRequests()
}

func handleRequests() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/search", search)
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
	dst, err := os.Create(path + ".pdf")
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	content, err := docconv.ConvertPath(path + ".pdf")
	if err != nil {
		log.Fatal(err)
	}
	var resume = Resume{Title: title, Path: path + ".pdf", Content: content.Body}
	db.Create(&resume)

	return c.String(http.StatusOK, "success")
}

func index(c echo.Context) error {
	resume := []Resume{}
	db.Find(&resume)

	return c.JSON(http.StatusOK, resume)
}

func search(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	resume := []Resume{}
	db.Where("content LIKE ?", "%"+keyword+"%").Find(&resume)

	return c.JSON(http.StatusOK, resume)
}
