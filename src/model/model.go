package model

import (
	"github.com/jinzhu/gorm"
)

type Resume struct {
	gorm.Model     `json:"-"`
	Email   string `json:"email" sql:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Path    string `json:"path"`
	Content string `json:"-" sql:"type:longtext CHARACTER SET utf8 COLLATE utf8_general_ci"`
}
