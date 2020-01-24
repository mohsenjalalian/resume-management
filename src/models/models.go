package models

import (
	"github.com/jinzhu/gorm"
)

type Resume struct {
	gorm.Model
	Title   string `json:"title" sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Path    string `json:"path"`
	Content string `json:"content" sql:"type:longtext CHARACTER SET utf8 COLLATE utf8_general_ci"`
}
