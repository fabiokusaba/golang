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