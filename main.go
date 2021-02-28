package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"p_accounting/models"
	_ "p_accounting/routers"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Tag{})
	if err != nil {
		fmt.Println(err)
		panic("failed migration")
	}

	err = db.AutoMigrate(&models.User{}, &models.Receipt{}, &models.Tag{})
	if err != nil {
		fmt.Println(err)
		panic("failed migration")
	}
}
