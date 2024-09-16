package main

import (
	"database/sql"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:WikitelecomuGr+dX@u2%@tcp(localhost:3306)/cursoGO")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = CreateTable(db)
	if err != nil {
		panic(err)
	}
	product := NewProduct("Celular", 1500.0)
	err = InsertProduct(db, *product)
	if err != nil {
		panic(err)
	}
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE Products (id varchar(100), name varchar(100), price decimal(10, 2));")
	if err != nil {
		panic(err)
	}
	return nil
}

func InsertProduct(db *sql.DB, product Product) error {
	query, err := db.Prepare("INSERT INTO products(id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}
