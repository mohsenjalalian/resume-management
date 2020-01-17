package models

import (
	"github.com/jinzhu/gorm"
)

type Resume struct {
	gorm.Model
	Title   string `json:"title"`
	Path    string `json:"path"`
	Content string `json:"content"`
}
