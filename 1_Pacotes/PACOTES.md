## Pacotes no Go
* Quando estamos lidando com mais de um pacote do mesmo projeto no Go a gente tem que criar uma coisa que chamamos de módulo.
* O módulo vai ser um conjunto de pacotes que compõe o projeto, então tanto pacotes que você criou como pacotes externos.
* Para criarmos um módulo é muito simples basta acessarmos a pasta que queremos fazer e digitar o comando: `go mod init <nome_do_modulo>`
* O módulo é o arquivo que vai centralizar todas as dependências do seu projeto.
* Quando rodamos o comando: `go build` ele criou um outro arquivo com o nome do nosso módulo, arquivo binário e executável, no qual ele
compilou todo o código presente no projeto.
* Como o pacote é um conjunto de arquivos que estão na mesma pasta se você quiser criar mais de um pacote você precisa criar mais de uma
pasta.
* Um detalhe muito importante como Go não é uma linguagem orientada a objetos não existe nada parecido com public, private, protected, o
que usamos para falar que determinada coisa como uma variável, uma função, uma struct, etc, são públicas é pela primeira letra delas, então
por exemplo se uma função começa com letra minúscula ela só vai ser visível dentro do pacote que ela está, se ela começa com letra maiúscula
ela vai poder ser exportada por outros pacotes.
* Dentro do próprio pacote eu consigo referenciar as funções que estão com letra minúscula, mas fora dele não.
* Uma coisa importante é que o nosso binário gerado, modulo, não se atualiza automaticamente então se você deseja atualizar ele é necessário
rodar novamente o comando `go build`.
* Outro comando do módulo seria o `go install` em que ele vai fazer a mesma coisa que o `go build`, porém ele vai salvar o arquivo em outro
lugar, ele vai salvar lá na raiz onde você instalou o Go.
* Para colocarmos um pacote externo para dentro do nosso módulo e assim conseguirmos usar ele em nosso projeto é bem simples: vamos até a raiz
na mesma pasta que está o nosso módulo e damos o comando: `go get <url_do_pacote>`. Quando usamos esse comando ele vai buscar esse pacote e
baixá-lo com todas as suas dependências, caso existam.
* O comando `go mod tidy` vai remover todas as dependências que não estão sendo usadas.
