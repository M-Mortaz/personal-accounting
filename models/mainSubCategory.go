package models

import "gorm.io/gorm"

type MainSubCategory struct {
	gorm.Model
	Name string
	SubCategory []SubCategory `gorm:"constant:onDelete:SET NULL;"`
}
