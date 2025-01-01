package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Tags do JSON -> por algum motivo você não quer que quando o seu cachorro for convertido para JSON você não quer que o campo
// Nome apareça, então você pode deixar ao invés da tag um traço "-" fazendo com que ele ignore esse campo porque ele não soube
// identificar no JSON, o mesmo vale também para o Marshal
type cachorro struct {
	Nome  string `json:"-"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// Declarando um objeto JSON
	cachorroEmJSON := `{"nome":"Rex", "raca":"Dálmata", "idade":3}`

	// Imagine agora que eu quero passar esse JSON, formato legível, para o nosso struct de cachorro, como faríamos esse processo?
	// Primeiramente precisamos criar a nossa variável do tipo cachorro sem passar valores a ela
	var c cachorro

	// Precisamos passar os nossos dados para dentro do struct cachorro para isso vamos chamar a função json.Unmarshal, ela recebe
	// dois parâmetros: o primeiro é os dados que queremos passar e o segundo parâmetro é o endereço de memória aonde a gente vai
	// jogar esses dados. E aqui passamos um ponteiro, endereço de memória, porque eu quero que essa variável fique alterada depois
	// dessa linha
	// Dois pontos importantes do Unmarshal: ele recebe no formato []byte e ele retorna só um valor que é um erro, então precisamos
	// tratar esse erro e para isso podemos utilizar o if init
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	// Exibindo o resultado em tela
	fmt.Println(c)

	// Convertendo para Map -> uma coisa que é interessante falar sobre o Map é que precisamos tomar muito cuidado com o tipo, imagine
	// que temos um map cuja as chaves são strings e o valor inteiro isso não vai dar erro logo de cara, mas quando a gente tentar converter
	// o nosso JSON para esse Map ele vai dar problema porque ele não vai conseguir converter uma string para int, ou seja, é um erro que só
	// é pego em execução, não é pego em compilação
	cachorro2EmJSON := `{"nome":"Toby", "raca":"Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
