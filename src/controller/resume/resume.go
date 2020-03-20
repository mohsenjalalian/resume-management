package resume

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
	"github.com/mohsenjalalian/resume-management/database/mysql"
	"github.com/mohsenjalalian/resume-management/model"
)

func New(c echo.Context) error {
	fl, err := c.FormFile("file")

	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("email")
	path := "statics/resumes/" + strconv.FormatInt(time.Now().UTC().Unix(), 10)

	// Source
	src, err := fl.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer src.Close()

	// Destination
	dst, err := os.Create(path + ".pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

	content, err := docconv.ConvertPath(path + ".pdf")
	if err != nil {
		log.Fatal(err)
	}
	var resume = model.Resume{Email: email, Path: path + ".pdf", Content: content.Body}
	mysql.Db.Create(&resume)

	return c.String(http.StatusOK, "success")
}

func Index(c echo.Context) error {
	resume := []model.Resume{}
	mysql.Db.Find(&resume)

	return c.JSON(http.StatusOK, resume)
}

func Search(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	resume := []model.Resume{}
	mysql.Db.Where("content LIKE ?", "%"+keyword+"%").Find(&resume)

	return c.JSON(http.StatusOK, resume)
}
