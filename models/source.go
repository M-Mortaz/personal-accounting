package models

import "gorm.io/gorm"

type Source struct {
	gorm.Model
	Name string
	Description string
	UserID uint
	Receipt []Receipt
}
