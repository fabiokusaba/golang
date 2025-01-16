package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// String de conexão MySQL
	stringConexao := "root:root123@/devbook?charset=utf8&parseTime=True&loc=Local"

	// Tentando abrir a conexão usando o pacote padrão do Go SQL
	// Chamando o método sql.Open passando dois parâmetros: o primeiro é uma string indicando o tipo de banco que estou me
	// referindo e o segundo parâmetro seria a string de conexão
	// Esse método retorna dois valores: o primeiro seria a conexão com o banco e o segundo seria um erro
	db, erro := sql.Open("mysql", stringConexao)
	// Tratando o erro
	if erro != nil {
		fmt.Println("Dentro do sql.Open")
		// Parando a execução do programa
		log.Fatal(erro)
	}

	// Fechando a conexão com o banco antes da função main terminar independente do que aconteça garantindo que ela não vai
	// permanecer aberta
	defer db.Close()

	// Testando a conexão com o método ping com if init
	// Perceba que aqui não utilizamos ´:=' porque estamos reaproveitando a variável erro
	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do db.Ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	// A partir desse ponto já estamos conectados no banco e já podemos começar a fazer operações nele
	// Para a gente fazer um select e retornar dados usamos o seguinte comando -> `db.Query`
	// Ele retorna dois valores pra gente: primeiro as linhas que ele achou da tabela e depois o erro
	linhas, erro := db.Query("select * from usuarios")
	// Tratando o erro
	if erro != nil {
		// Existem diversas maneiras de tratarmos esse erro mas estamos usando um log.Fatal para que ele pare a execução de
		// tudo
		log.Fatal(erro)
	}

	// Da mesma forma que fizemos com o banco de dados as linhas a gente pode usar `linhas.Close` para fechá-las, porque
	// elas são uma espécie de cursor que vai apontando para a próxima linha dentro da tabela e pra gente fazer a leitura a
	// gente vai iterando por esse cursor, então sempre fechamos quando terminamos a execução mesma coisa que faríamos com o
	// banco pra não deixar esse cara aberto
	defer linhas.Close()

	fmt.Println(linhas)
}
