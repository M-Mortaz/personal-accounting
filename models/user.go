package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"index:,unique"`
	FirstName   string
	LastName    string
	Email       string `gorm:"index:,unique"`
	PhoneNumber string `gorm:"index:,unique"`
	Password    string
	Active      bool
	SuperUser   bool
	LastSignup  time.Time
	Receipts    []Receipt
	Source []Source
	Tag []Tag
	Category []Category
	SubCategory []SubCategory
}
