package models

import "gorm.io/gorm"

type Receipt struct {
	gorm.Model
	UserID uint
	TagID uint
}
