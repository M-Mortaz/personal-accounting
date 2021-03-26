package models

import "gorm.io/gorm"

type SubCategory struct {
	gorm.Model
	Name string
	CategoryID uint
	UserID uint
	MainSubCategoryID uint
	Receipt []Receipt
}
