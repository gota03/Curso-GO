package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"` // Informa ao GORM que este campo é a chave primária.
	Name  string
	Price float64
}

func main() {
	dsn := "root:WikitelecomuGr+dX@u2%@tcp(localhost:3306)/cursoGO"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Abre uma conexão com o banco de dados usando a string de conexão dsn e o driver MySQL.

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Product{})
	// Cria ou atualiza a tabela no banco de dados baseada na struct Go. Ele mapeia automaticamente os campos da struct para as colunas da tabela.
	if err != nil {
		panic(err)
	}

	// Inserindo um registro único
	db.Create(Product{
		Name:  "Celular",
		Price: 2500.0,
	})
	// Insere novos registros no banco de dados. Essa função pode ser usada tanto para inserir um único registro quanto vários registros de uma vez
	// Se um campo como ID for uma chave primária autoincrementada, o GORM gerará automaticamente os valores corretos para ele.

	// Inserindo múltiplos registros
	// products := []Product{
	// 	{Name: "TV", Price: 3000.0},
	// 	{Name: "Iphone", Price: 4500.0},
	// 	{Name: "Carregador", Price: 150.0},
	// 	{Name: "Fone de ouvido", Price: 300.0},
	// }
	// db.Create(&products)

	// Fazendo uma busca única por id
	var product Product
	db.First(&product, 1)
	db.First(&product, "name = ?", "Iphone")

	// Fazendo múltiplas buscas
	var products []Product
	db.Find(&products)
	for _, v := range products {
		fmt.Printf("ID: %v\nName: %v\nPrice: %v\n", v.ID, v.Name, v.Price)
	}
}
