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
	ID          int `gorm:"primaryKey"`
	Name        string
	Price       float64
	Category_ID int      //  É o campo numérico que armazena o valor da chave estrangeira (ID da categoria) no banco de dados. Isso é o que define o relacionamento de fato no nível do banco de dados.
	Category    Category // Esse campo permite que o GORM carregue e associe automaticamente os detalhes completos da categoria ao produto. Quando você usa Preload("Category"), o GORM faz uma consulta adicional para buscar os dados completos da categoria e preencher essa struct.

	// O GORM usa o valor de Category_ID (por exemplo, 1) para encontrar o registro correspondente na tabela categories.
	// Ele então preenche o campo Category da struct Product com a instância da Category que tem ID igual a Category_ID.

	gorm.Model
}

func main() {
	dsn := "root:WikitelecomuGr+dX@u2%@tcp(localhost:3306)/cursoGO?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Products{}, &Category{})
	var products []Products
	db.Preload("Category").Find(&products)
	// Preload carrega os dados relacionados da tabela categories para cada registro encontrado na tabela products. O GORM faz isso em uma única consulta SQL usando JOIN ou múltiplas consultas internas, dependendo do contexto.
	// Garante que os dados relacionados da tabela categories sejam carregados junto com os produtos

	for _, v := range products {
		fmt.Printf("ID: %v\nName: %v\nPrice: %.2f\nCategory: %v\n", v.ID, v.Name, v.Price, v.Category.Name)
		// Se v.Category_ID é 1, o GORM carrega a categoria com ID 1 da tabela categories, que é "Eletrônicos".
		// Ele então associa essa categoria ao produto e você pode acessar v.Category.Name para obter o nome da categoria.
	}
}
