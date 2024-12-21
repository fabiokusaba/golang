# Operadores

## Aritméticos
* Operações matemáticas
* `+` - Adição
* `-` - Subtração
* `/` - Divisão
* `*` - Multiplicação
* `%` - Módulo

* Quando temos duas variáveis declaradas uma do tipo `int16` e outra do tipo `int32` e tentamos somá-las perceba que isso
não vai funcionar porque você não pode somar, subtrair, comparar, você não pode fazer nada no Go com variáveis que são de
tipos diferentes, ou seja, em Go só é permitido realizar operações com tipos iguais.

## Atribuição
* Existem dois jeitos de você atribuir valor a uma variável: o primeiro deles é o `=` ou você pode usar o `:=`.

## Relacionais
* `==` - Igual
* `>` - Maior
* `>=` - Maior ou Igual
* `<` - Menor
* `<=` - Menor ou Igual
* `!=` - Diferente

* Esses operadores sempre retornam um valor booleano `true` ou `false`.

## Lógicos
* `&&` - Retorna `true` se todas as condições forem avaliadas como `true`
* `||` - Retorna `true` se uma das condições for avaliada como `true`
* `!` - Negar, ou seja, é você inverter o valor de uma variável booleana

## Unários
* `++`
* `--`

* Operadores que interagem com uma variável por vez, por exemplo quando você quer incrementar o valor de uma variável.
* Em algumas linguagens você pode se deparar com o incremento e decremento pré e pós fixado, ou seja, você realizar a
operação antes ou depois do número ser avaliado, no Go isso não é possível.

## Ternário
* Operador ternário ele é um operador que muitas linguagens usam pra você colocar um valor numa variável, em geral, baseado
em uma condição, não existe isso em Go.
* Resumindo o operador ternário tem uma condição a ser avaliada, um valor caso a condição seja verdadeira e um valor para a
variável caso a condição seja falsa.
* Isso não existe em Go porque a premissa do Go é que exista só uma maneira de fazer cada coisa, a exceção poderia ser a
declaração de variáveis que como vimos existe mais de uma forma de se fazer.
* Então, em Go usaríamos o `if-else` nesse caso.