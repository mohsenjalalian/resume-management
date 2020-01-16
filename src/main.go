package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"code.sajari.com/docconv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/mohsenjalalian/resume-management/database"
	"github.com/mohsenjalalian/resume-management/migration"
	"github.com/mohsenjalalian/resume-management/models"
)

func main() {
	mysql.Open()
	if mysql.Err != nil {
		log.Println("Connection Failed to Open")
	}
	migration.Migrate()
	defer mysql.Db.Close()
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
	var resume = models.Resume{Title: title, Path: path + ".pdf", Content: content.Body}
	mysql.Db.Create(&resume)

	return c.String(http.StatusOK, "success")
}

func index(c echo.Context) error {
	resume := []models.Resume{}
	mysql.Db.Find(&resume)

	return c.JSON(http.StatusOK, resume)
}

func search(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	resume := []models.Resume{}
	mysql.Db.Where("content LIKE ?", "%"+keyword+"%").Find(&resume)

	return c.JSON(http.StatusOK, resume)
}
