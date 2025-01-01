## Testes Automatizados
* De uma forma bem resumida ele vai ser uma função que vai testar outra função sua e ver se o resultado dela é o que você está
esperando realmente.
* Em grandes aplicações é um cenário bastante comum você estar criando testes para que você garanta o comportamento das coisas
então imagine que você tem uma função que recebendo dois parâmetros ela tem que retornar um resultado específico, os testes são
basicamente para garantir que a sua função recebendo esses parâmetros vai, de fato, retornar o resultado que você está esperando.
* É um jeito de você garantir que o que você implementou está certo e que ele também vai continuar certo ao longo do tempo.
* A vantagem do teste é que ele te dá uma segurança muito grande para fazer mudanças no seu código, imagine que hoje eu tenho uma
função que está funcionando, retornando o resultado que eu espero que ela retorne, amanhã faço uma alteração nela, se ela parar de
devolver o resultado que estou esperando por conta dessa alteração meu teste vai acusar isso e dessa forma você pode avaliar se a
alteração realmente faz sentido ou se é algum efeito colateral que você não tinha previsto.
* Para fazermos os testes vamos utilizar um pacote no Go chamando testing e para que tudo funcione perfeitamente existem algumas
condições a serem seguidas: o teste de uma função nunca fica no mesmo arquivo da função em si, isso porque os arquivos de teste
para serem reconhecidos pelo Go eles tem que ter um nome específico `<nome_test.go>`.
* Para rodar todos os nossos testes vamos fazer pela linha de comando então chamamos um comando específico do Go que vai entrar nos
arquivos que `_test.go`, que ele entende ser um arquivo de teste, e vai começar a executar as funções de teste dentro dele.
* Teste de Unidade: é um teste que vai testar a menor unidade do seu código.
* O teste é a única exceção em que você pode ter dois pacotes diferentes dentro da mesma pasta.
* Quando você roda o comando `go test` por padrão ele vai executar os testes seguindo as convenções que falamos do pacote que você se
encontra existe uma maneira dele executar os testes que estão em outros pacotes dentro do seu projeto para isso utilizamos o comando
`go test ./...`.
* Quando não fazemos nenhuma alteração desde a última vez que rodamos o comando `go test` o Go é esperto o suficiente para não rodar
o comando novamente, então ele vai ficar em cached porque ele sabe que aquela função que foi testada já passou.
* Quando utilizamos a flag `-v`, significa verbose, é o modo em que ele vai escrever o teste de forma mais detalhada na tela.
* É possível rodar os testes em paralelo, para isso precisamos colocar `t.Parallel()`. Colocar esse trecho de código em apenas uma função
não adianta nada, é preciso colocar em outras funções também pois ele só vai rodar em paralelo as funções que tiverem marcadas com esse cara
então necessariamente é preciso colocá-lo logo no início, na primeira linha.
* Como você vê a porcentagem que está sendo coberto pelo seu teste? Podemos rodar o seguinte comando `go test --cover`.
* Existem algumas ferramentas no Go que nos auxiliam a checar a taxa de cobertura dos nossos testes uma delas é a geração de um arquivo txt
que vai conter um relatório das linhas que estão cobertas, para fazer isso é bem simples `go test --coverprofile <nome_arquivo.txt>`.
* Podemos ler esse arquivo que geramos com o comando acima de duas maneiras: podemos usar o comando `go tool cover --func=<nome_arquivo.txt>`.
Então, ele vai ler o arquivo txt pra gente, vai entender o que o arquivo quer dizer e exibir no terminal.
* Para descobrirmos que linha em específico não está coberta pelos testes podemos usar o comando `go tool cover --html=<nome_arquivo.txt>`.
Ele vai mostrar pra gente um arquivo html que vai ter visualmente todas as linhas que não estão cobertas.
* Como mencionado no começo do curso você só pode ter um pacote por pasta dentro do Go, a exceção seria para você fazer os testes, então você
poderia fazer assim: `package enderecos_test`, ou seja, o mesmo nome do pacote que você está adicionando `_test` ao final. Quando optamos por
essa nomenclatura de pacote precisamos fazer o import do pacote que queremos usar, por exemplo `import "introducao-testes/enderecos"`, ou seja,
usaríamos o nome do módulo e do pacote. 
* O que muitas vezes é feito quando usamos essa sintaxe é você importar esse pacote com um alias, ou seja, um apelido pra você não precisar
ficar fazendo `enderecos.TipoDeEndereco()`, assim: `import . "introducao-testes/enderecos"`. Quando você faz isso o Go vai entender que o alias
vai se referir ao pacote por padrão.