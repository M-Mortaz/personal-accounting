package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"p_accounting/models"
	_ "p_accounting/routers"
)

func main() {
	dsn := "host=localhost user=p_acc password=p_acc dbname=p_acc"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&models.MainCategory{})
	err = db.AutoMigrate(
		&models.User{},
		&models.Tag{},
		&models.Source{},
		&models.MainCategory{},
		&models.MainSubCategory{},
		&models.Category{},
		&models.SubCategory{},
		&models.Receipt{},
	)
	if err != nil {
		fmt.Println(err)
		panic("failed migration")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("http://localhost:8000/")
	err = r.Run(":8000")
	if err != nil {
		panic("Some issue occurs ^-^")
	}
}
