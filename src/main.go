package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"	
)

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
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
