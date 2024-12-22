## Switch
* Para usarmos o switch é bem simples, usamos a cláusula `switch` e temos duas formas de se utilizar: uma é passando diretamente
uma condição, variável, alguma coisa para ser avaliada, ou deixar essa condição para ser avaliada dentro de cada caso.
* Para demonstrarmos o uso do switch criamos uma função diaDaSemana que nos retorna uma string correspondente ao dia da semana com
base no número que passamos, perceba que se o número que for passado estiver entre 1 a 7 a função vai nos retornar um de seus casos
mas se por acaso esse número não estiver entre 1 e 7 ele não vai cair em nenhuma das condições do nosso switch e o Go vai nos alertar
quanto a isso e não vai deixar isso acontecer, então para contornarmos isso com o switch depois de todos os nossos casos podemos usar
a cláusula `default` para atribuir um valor padrão caso nenhuma das condições seja atendida.
* Existe uma cláusula que não é muito utilizada no Go, somente para alguns padrões muito específicos que você usa dentro do switch, que
é a cláusula chamada `fallthrough`, o que acontece é que você está falando o seguinte: ele vai avaliar essa condição e vai executar o que
estiver ali dentro essa cláusula `fallthrough` ele vai jogar o seu código para dentro da próxima condição.
* Importante frisar que no Go você não precisa usar a cláusula `break` que existe em muitas linguagens de programação, o break faz o
seguinte: quando o seu switch avalia uma condição como verdadeira você precisa dizer a ele o que fazer e se você quiser que ele não execute
as outras condições você precisa usar o break que vai sinalizar para ele que ele precisa sair do switch, no Go não existe essa cláusula
porque ele já sai automaticamente, então por padrão quando você avalia uma condição como verdadeira o Go vai executar esse trecho de código
e vai ignorar todas as outras.