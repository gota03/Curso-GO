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
	gorm.Model
	// gorm.Model herda automaticamente campos que ajudam a gerenciar timestamps e exclusões lógicas
	// Quando chama a função Delete do GORM, em vez de remover o registro fisicamente do banco de dados, ele apenas preenche o campo DeletedAt com a data da exclusão. Isso permite que mantenha um histórico de registros "excluídos", sem removê-los permanentemente.
}

func main() {
	dsn := "root:WikitelecomuGr+dX@u2%@tcp(localhost:3306)/cursoGO?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Abre uma conexão com o banco de dados usando a string de conexão dsn(data sour	ce name) e o driver MySQL.

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Product{})
	// Cria ou atualiza a tabela no banco de dados baseada na struct Go. Ele mapeia automaticamente os campos da struct para as colunas da tabela.
	// GORM precisa da estrutura completa para entender como mapear os campos da struct para as colunas no banco de dados. Usando um ponteiro, ele consegue trabalhar diretamente com a instância, garantindo que a função possa modificar ou inspecionar os valores dos campos da struct, o que é útil para realizar a migração corretamente.

	if err != nil {
		panic(err)
	}

	// Inserindo um registro único
	db.Create(&Product{
		Name:  "Celular",
		Price: 2500.0,
	})
	// O uso de um ponteiro é necessário porque ele pode modificar a struct durante a operação (como, por exemplo, preencher o campo ID com o valor gerado automaticamente pelo banco de dados). Sem o ponteiro, a função não poderia alterar o valor original da struct, o que é importante para que você possa reutilizar a struct após a operação.
	// Insere novos registros no banco de dados. Essa função pode ser usada tanto para inserir um único registro quanto vários registros de uma vez
	// Se um campo como ID for uma chave primária autoincrementada, o GORM gerará automaticamente os valores corretos para ele.

	// Inserindo múltiplos registros
	products := []Product{
		{Name: "TV", Price: 3000.0},
		{Name: "Iphone", Price: 4500.0},
		{Name: "Carregador", Price: 150.0},
		{Name: "Fone de ouvido", Price: 300.0},
	}
	db.Create(&products)

	// Fazendo uma busca única por id
	var product Product
	db.First(&product, 2)
	// O GORM assume que o segundo argumento (neste caso, 2) é a chave primária da tabela (por padrão, o campo ID)
	// O resultado da consulta é armazenado na variável product, que é passada como um ponteiro (&product). Isso permite que o GORM preencha os campos da struct com os valores do registro encontrado no banco de dados.

	db.First(&product, "name = ?", "Iphone")
	// usa um filtro condicional no campo name e o valor "Iphone" substitui o ?.

	// Fazendo múltiplas buscas
	var products2 []Product
	db.Find(&products2)
	// Retorna todos os registros que atendem aos critérios fornecidos. Se nenhum critério for fornecido, ele busca todos os registros da tabela.
	// É passado um ponteiro para o slice para que ele possa ser preenchido com os dados da consulta.

	for _, v := range products {
		fmt.Printf("ID: %v\nName: %v\nPrice: %v\n", v.ID, v.Name, v.Price)
	}

	var products3 []Product
	db.Limit(2).Offset(1).Find(&products3)
	// A função Limit define quantos registros serão retornados pela consulta
	// Offset define a partir de qual posição os registros serão buscados, Offset(1) faz com que a consulta ignore o primeiro registro da tabela e comece a partir do segundo.

	for _, v := range products3 {
		fmt.Printf("%v", v)
	}

	var products4 []Product
	db.Where("price > ?", 400).Find(&products4)
	// Where no GORM é utilizada para definir condições (filtros) em consultas SQL. Ela adiciona cláusulas WHERE na consulta para buscar apenas os registros que atendem aos critérios especificados.
	// fmt.Printf("%v", products)

	var p Product
	db.First(&p, 1)
	p.Price = 2300.0
	db.Save(&p)
	// Save no GORM é utilizada para salvar ou atualizar um registro no banco de dados. Quando altera um valor de um campo em uma struct e chama Save, o GORM gera uma consulta UPDATE para persistir essas mudanças no banco.

	var p2 Product
	db.First(&p2, 1)
	db.Delete(&p2)
	// A função Delete no GORM é usada para excluir um registro do banco de dados. Ela gera uma consulta SQL DELETE para remover o registro correspondente.
}
