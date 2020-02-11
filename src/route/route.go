package route

import (
	"github.com/labstack/echo/v4"
	"github.com/mohsenjalalian/resume-management/controller/resume"
)

func HandleRequests() {
	e := echo.New()
	e.GET("/", resume.Index)
	e.GET("/search", resume.Search)
	e.POST("/", resume.New)
	e.Static("/statics", "statics")
	e.Logger.Fatal(e.Start(":5600"))
}
