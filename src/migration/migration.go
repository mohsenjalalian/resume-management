package migration

import (
	"github.com/mohsenjalalian/resume-management/database/mysql"
	"github.com/mohsenjalalian/resume-management/model"
)

func Migrate() {
	mysql.Db.AutoMigrate(&model.Resume{})
}
