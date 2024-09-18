package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Products struct {
	ID             int `gorm:"primaryKey"`
	Name           string
	Price          float64
	CategoryID     int
	Category       Category
	SerialNumberID int
	SerialNumber   SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:WikitelecomuGr+dX@u2%@tcp(localhost:3306)/cursoGO?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)

	}
	db.AutoMigrate(&Category{}, &Products{}, &SerialNumber{})
	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	db.Create(&SerialNumber{
		Number:    "12345",
		ProductID: 1,
	})

	db.Create(&Products{
		Name:           "Iphone",
		Price:          4200.0,
		CategoryID:     1,
		SerialNumberID: 1,
	})
	var products []Products
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Printf("ID: %v\nName: %v\nPrice: %.2f\nCategory: %v\nSerialNumber: %v", product.ID, product.Name, product.Price, product.Category.Name, product.SerialNumber.Number)
	}
}
