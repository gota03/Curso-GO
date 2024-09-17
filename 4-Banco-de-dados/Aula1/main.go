package main

import (
	"database/sql"
	"fmt"

	// O pacote sql fornece uma interface genérica para interagir com diferentes bancos de dados, neste caso, o MySQL.

	_ "github.com/go-sql-driver/mysql"
	// Este é o driver específico para MySQL, necessário para que Go possa interagir com o banco de dados MySQL.
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:WikitelecomuGr+dX@u2%@tcp(localhost:3306)/cursoGO")
	// Abre uma conexão com o banco de dados
	// 1º parâmetro: especifica o driver MySQL.
	// 2º parâmetro: especifica as credenciais de conexão, incluindo o usuário (root), a senha e o endereço do servidor (localhost:3306) e o nome do banco (cursoGO).

	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Fecha a conexão com o banco de dados

	err = CreateTable(db)
	if err != nil {
		panic(err)
	}

	product := NewProduct("Carregador", 150.0)
	err = InsertProduct(db, product)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Produto inserido com ID: %d\n", product.ID)

	product.Price = 100.0
	product.Name = "Carregador iphone"
	err = UpdateProduct(db, product)
	if err != nil {
		panic(err)
	}
	product, err = SelectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %v\nName: %v\nPrice: %.2f", product.ID, product.Name, product.Price)
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE products (id INT AUTO_INCREMENT PRIMARY KEY, name varchar(100), price decimal(10, 2));")
	// Executa uma consulta SQL
	if err != nil {
		panic(err)
	}
	return nil
}

func InsertProduct(db *sql.DB, product *Product) error {
	query, err := db.Prepare("INSERT INTO products(name, price) VALUES (?, ?);")
	// prepara uma instrução SQL para ser executada várias vezes com diferentes valores. Ela é útil quando deseja executar a mesma consulta repetidamente, trocando apenas os parâmetros.
	// é mais eficiente para consultas repetidas, porque a consulta SQL é analisada (parsed) apenas uma vez pelo banco de dados, e os parâmetros são fornecidos posteriormente nas execuções.
	// Os pontos de interrogação (?) na string SQL de db.Prepare são placeholders que indicam onde os valores dos parâmetros serão inseridos.

	if err != nil {
		return err
	}
	defer query.Close()

	result, err := query.Exec(product.Name, product.Price)
	// executa diretamente uma instrução SQL (como INSERT, UPDATE, ou DELETE), sem a necessidade de prepará-la previamente. Essa função é útil para consultas que você não precisa reutilizar.
	// os argumentos fornecidos são mapeados diretamente, na ordem, para os placeholders na string do argumento da função db.Prepare.

	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()
	// retorna o inteiro gerado pelo banco de dados em resposta a um comando. Normalmente, isso ocorrerá em uma coluna de "incremento automático" ao inserir uma nova linha.

	if err != nil {
		return err
	}

	product.ID = int(lastInsertId)
	fmt.Printf("ID: %v\nName: %v\nPrice: %v\n", product.ID, product.Name, product.Price)
	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? where id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		println("Nenhuma linha foi atualizada, verifique o ID")
		return nil
	}

	fmt.Printf("Atualizando produto:\nID:%v\nName:%v\nPrice:%.2f\n", product.ID, product.Name, product.Price)

	return nil
}

func SelectProduct(db *sql.DB, id int) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price from products where id = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	// QueryRow:
	// Executa a consulta preparada e espera que ela retorne apenas uma linha do resultado.
	// O argumento id fornecido aqui é usado para preencher o placeholder (?) da consulta SQL, ou seja, ele substitui o ? pelo valor do id passado para a função.
	// Se a consulta retornar mais de uma linha, QueryRow irá selecionar apenas a primeira.
	// Se a consulta não encontrar nenhuma linha correspondente, o erro sql.ErrNoRows será retornado.

	// Scan:
	// Extrai os valores das colunas retornadas e os copia para as variáveis fornecidas. As variáveis devem ser passadas por referência (usando &), e a ordem e quantidade devem corresponder às colunas da consulta SQL.
	// Se a consulta não encontrar nenhuma linha, Scan retornará um erro (sql.ErrNoRows).

	if err != nil {
		return nil, err
	}
	return &p, nil
}
