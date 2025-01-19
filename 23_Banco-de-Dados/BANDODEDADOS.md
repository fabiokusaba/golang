## Banco de Dados - MySQL
### Comandos básicos
* Criando um banco de dados -> `CREATE DATABASE <nome_do_banco>;`
* Selecionando o banco de dados -> `USE <nome_do_database>;`, através desse comando fazemos com que todos os demais comandos
que rodarmos sejam aplicados a esse banco em específico.
* Criando uma tabela no banco de dados -> `CREATE TABLE <nome_da_tabela> (<colunas>);`
* Mostrar as tabelas do banco de dados -> `SHOW TABLES;`
* Para abrirmos a conexão com o banco de dados vamos precisar fazer o download de um pacote que é o driver do MySQL, banco de
dados que estamos utilizando, para isso usamos o seguinte comando -> `go get github.com/go-sql-driver/mysql`.
* O SQL é um dialeto que é usado por muitos bancos de dados, então você tem MySQL, Postgres, e cada um desses caras tem um 
driver específico pra funcionar junto com o pacote SQL que é padrão do Go.
* Para fazer a conexão pelo Go vamos precisar usar uma string de conexão que vai passar o usuário, senha e o nome do banco que
a gente quer conectar. Uma estrutura básica seria -> `usuario:senha@/banco`, podemos passar alguns parâmetros adicionais como:
`?charset=utf8`, reconhecer caracteres com maior facilidade, `&parseTime=True`, para campos de data, `&loc=Local`, configurar
o locale.
* O Go não deixa você importar um pacote e não usar, o que fazemos com import do driver do MySQL é chamado de import implícito
ou seja, a gente vai usar ele mas não especificamente nesse arquivo que acabamos de criar e para fazer isso a gente usa um `_`
na frente do nosso import.
* Não vai ser o nosso arquivo `banco-de-dados.go` que vai usar o driver do MySQL, quando passamos uma string `"mysql"` estamos
dizendo que a conexão vai precisar de um driver MySQL, o que vai acontecer é que o pacote sql do Go vai saber que a conexão é
com MySQL e ele é quem vai procurar o driver e vai utilizá-lo, então eu importo o pacote do driver aqui mas quem vai usar vai
ser o pacote sql por isso importamos com o `_` na frente.