package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"crud/database"

	"github.com/gorilla/mux"
)

// Podemos criar com a inicial do nome em minúsculo porque só iremos utilizá-lo aqui dentro desse pacote, irá conter todos os
// campos da nossa tabela
type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// Primeiro passo é ler o corpo da requisição
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		// Caso não tenha o corpo da requisição meu request já para de funcionar, não tem como prosseguir a inserção de um usuário
		// no banco de dados sem os seus dados
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		// Mesmo que a sua função não retorne valores você pode usar um return aqui sem especificar nada depois dele só para que ele
		// pare a execução
		return
	}

	// Com essa requisição eu vou jogar ela dentro de um struct de usuário, como a requisição que a gente mandou é em json podemos
	// usar o json.Unmarshal
	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	// Para inserir o nosso usuário no banco precisamos primeiro abrir a conexão
	db, erro := database.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	// Fechando a conexão com o banco ao final
	defer db.Close()

	// Agora para inserir um usuário no banco utilizando o terminal do MySQL seria: INSERT INTO usuarios (nome, email) VALUES(?, ?),
	// podemos fazer um processo muito parecido no Go para isso vamos utilizar prepared statement, com isso você cria um comando de
	// inserção que é muito utilizado para evitar ataques de sql injection, no prepared statement criamos o comando mas não passamos
	// os valores porque passando os valores diretamente podemos sofrer ataques de sql injection com a manipulação dos valores
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	// O statement da mesma forma que o banco precisa ser fechado depois então podemos usar um defer statement.Close
	defer statement.Close()

	// Uma vez que o nosso statement está construído precisamos passar os valores para ele e para isso usamos a função Exec respeitando
	// a ordem
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	// Se eu passar para essa linha significa que o usuário já foi inserido no banco e podemos retornar o id desse usuário
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	// No Go costumamos passar o status code no w.WriteHeader
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// Diferente da função de CriarUsuario não vai vir nenhum corpo na requisição, vou apenas realizar uma consulta no banco
	// A primeira coisa que vou fazer é abrir a conexão com o banco de dados
	db, erro := database.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	// Fechando a conexão com o banco ao final
	defer db.Close()

	// Aqui já posso executar o comando que vai trazer todos os usuários -> SELECT * FROM usuarios
	// Quando vamos fazer consultar usamos outro comando que é o Query que vai retornar linhas da tabela pra gente
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	// Da mesma forma que o statement a gente também precisa fechar esse cara
	defer linhas.Close()

	// Aqui vou criar um slice de usuarios
	var usuarios []usuario

	// Como a minha consulta vai trazer todos os usuários ele pode retornar mais de uma linha da tabela então vou ter
	// que iterar por essas linhas e montar vários structs de usuario a medida que elas chegam
	for linhas.Next() {
		// Para cada iteração vou criar um usuario, ler os dados que estão nessa linha e jogá-los para dentro desse struct
		// e depois jogar para dentro do meu slice de usuarios
		var usuario usuario

		// Para fazer a leitura dos dados vou utilizar o método Scan, então vou scanear cada uma das informações dessa linha
		// na ordem que elas aparecem no banco e vou jogar nas propriedades do nosso usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		// Se der tudo certo com o nosso Scan a gente já vai ter um usuário aqui com todos os campos populados então podemos
		// jogar esse usuario no nosso slice de usuarios
		usuarios = append(usuarios, usuario)
	}

	// Saindo do nosso for quer dizer que todos os dados foram lidos e podemos retorná-los
	w.WriteHeader(http.StatusOK)
	// Nesse caso preciso transformar esse slice de usuarios em um json e para isso vamos usar a função json.NewEncoder
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para json"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// Primeiramente vamos ler o parâmetro que está sendo passado na rota e para isso vamos utilizar o gorilla mux
	// Uma requisição pode ter vários parâmetros então aqui vamos retornar todos eles, vou conseguir identificar o
	// parâmetro que queremos pegar baseado no nome que passamos na uri
	parametros := mux.Vars(r)

	// O parâmetro vai vir como uma string então precisamos converter para um número inteiro, essa função ParseUint
	// recebe três parâmetros: o que queremos converter, a base do número (10 - decimal), tamanho dos bits (32)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	// Depois desse ponto eu já tenho o ID e posso abrir a conexão com o banco
	db, erro := database.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	// Fechando a conexão com o banco ao final
	defer db.Close()

	// Executando a consulta que atenda a condição do ID que foi passado
	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	// Depois desse cara eu vou criar um usuario
	var usuario usuario

	if linha.Next() {
		// Executando o Scan e jogando para dentro do nosso usuario
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	// Fazendo o Encoder dos dados
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para json"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// Da mesma forma que fizemos no BuscarUsuario vamos começar primeiramente lendo o parâmetro
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	// Agora vou ler o corpo da requisição para depois abrir a conexão com o banco, seguindo a sugestão de só abrir a conexão com
	// o banco depois de todos os pré-requisitos da atualização estejam preenchidos
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	// A requisição está em formato json então vou passar ela para o meu struct que vou criar
	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	// Nesse ponto, se nada der erro, temos os requisitos preenchidos para abrirmos a conexão com o banco
	db, erro := database.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	// Fechando a conexão ao final
	defer db.Close()

	// Agora o próximo passo é fazer o statement, sempre usamos o statement para qualquer operação que não seja de consulta
	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	// Fechando o statement ao final
	defer statement.Close()

	// Executando o statement
	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	// Primeiramente vamos começar lendo os parâmetros
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	// Não vai ter nada no corpo da requisição do delete então a partir desse ponto eu posso abrir a conexão com o banco
	db, erro := database.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	// Fechando a conexão ao final
	defer db.Close()

	// Vamos utilizar o prepared statement
	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	// Fechando o statement ao final
	defer statement.Close()

	// Para finalizar vamos executar esse statement através da função Exec
	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
