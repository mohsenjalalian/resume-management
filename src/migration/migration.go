package migration

import (
	"github.com/mohsenjalalian/resume-management/database"
	"github.com/mohsenjalalian/resume-management/models"
)

func Migrate() {
	mysql.Db.AutoMigrate(&models.Resume{})
}
