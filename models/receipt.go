package models

import (
	"gorm.io/gorm"
	"time"
)

type Receipt struct {
	gorm.Model
	Value int
	Description string
	HappenAt time.Time
	Hidden bool
	Type string
	UserID uint
	TagID uint
	SourceID uint
	SubCategoryID uint
}
