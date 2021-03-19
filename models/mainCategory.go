package models

import "gorm.io/gorm"

type MainCategory struct {
	gorm.Model
	Name string
	Category []Category `gorm:"constant:onDelete:SET NULL;"`
}
