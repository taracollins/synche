package schema

import "gorm.io/gorm"

type Directory struct {
	gorm.Model
	Path string `gorm:"unique"`
}

