package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

func Conectar() (*sql.DB, error) {
	stringConexao := "root:root123@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		// Nesse caso precisamos dar um retorno um pouco diferente, não posso dar um log.Fatal aqui porque diferente da função main
		// que estávamos usando nas outras vezes essa função aqui tem retorno, então sempre que eu for cair em algum error e eu precisar
		// sair da função eu tenho que dar um return, nesse caso preciso retornar esses dois caras: *sql.DB, error, só que via de regra
		// você só vai ter um dos dois porque eles são mutuamente exclusivos, ou seja, se você tiver o banco é porque não deu error, se
		// você tem o error é porque você não tem o banco, quando isso acontecer precisamos fazer da seguinte forma:
		return nil, erro
	}

	// Nesse caso aqui eu não vou dar um defer db.Close porque é uma função externa que vai ser chamada pelas minhas funções que vão
	// tratar requisição http e eu preciso que a conexão chegue para elas aberta, vai ser dentro das funções que vão trabalhar com o
	// http que eu vou dar o defer db.Close
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
