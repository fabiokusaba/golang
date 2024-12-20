# Tipos de Dados

## Números
### Inteiros
* Temos dois tipos: os inteiros e os reais, ou seja, int e float só que eles se subdividem em várias categorias.
* Existem quatro tipos de números inteiros em Go: você tem o `int8`, `int16`, `int32` e o `int64`. A diferença desses
caras é a quantidade de bits que eles ocupam, assim o `int8` vai suportar um número de até 8 bits, o `int16` até 16 bits
então quanto maior for o seu int maior será a capacidade que ele vai suportar, por exemplo o `int64` vai suportar números
enormes.
* Além desses quatro tipos de int que demonstramos existe também o `int` sozinho, quanto você não especifica o valor
dele ele vai usar a arquitetura do seu computador como base, então por exemplo se o seu computador for 64 bits ele vai 
criar um `int64`, se o seu computador for 32 bits ele vai criar um `int32`, é uma forma que você não precisa especificar e
ele vai usar de acordo com a sua arquitetura.
* Importante lembrar que esses ints suportam números negativos sem problema nenhum isso é relevante porque além do int 
dentro do Go a gente tem também o `uint` que quer dizer "unsigned int" que é um int sem sinal e também segue a mesma 
convenção do int na questão dos bits.
* Interessante que esses tipos de dados no Go alguns deles tem uma coisa chamada alias que é um apelido, então desses 
números inteiros dois deles podemos usar apelidos para nos referirmos. No caso do inteiro que aceita números negativos se
você quiser declarar um `int32`, por exemplo, você pode declarar dessa forma `int32` ou você pode colocar `rune`.
* E também tem um alias para o `uint8` onde você pode colocar `byte`.

### Reais
* São mais fáceis de entender porque só possuem dois tipos: `float32` e o `float64`.
* Seguem a mesma lógica, então `float32` porque são 32 bits, `float64` são 64 bits, a diferença agora é que esses números
suportam vírgula.
* Você não pode declarar explicitamente só `float` você tem que declarar o `float32` ou `float64`.
* Se você for pela inferência ele vai colocar como `float` e vai ser o mesmo esquema da arquitetura, se for um arquitetura
de 32 bits ele vai pegar o `float32` se for de 64 bits ele vai pegar o `float64`.

## Strings
* É uma cadeia de caracteres.
* Uma coisa importante de se ressaltar é que no Go não temos o tipo `char` que seria um caracter só.
* O que seria `char` em Go vai ser convertido para um número.
* Importante frisar que no Go sempre utilize aspas duplas `""` para `string`, aspas simples `''` ele vai considerar o que
seria o `char`.
* O mais perto do `char` que temos no Go é: se você declarar algo com aspas simples `'B'`, um caracter só, ele vai pegar o
número da tabela ASCII desse caracter e se você reparar bem esse cara aqui é um `int` porque ele vai sempre referenciar um
número, então não existe o tipo de dados `char` em Go.

## Valor Zero
* É importante falarmos de um conceito da linguagem que é o valor zero, o valor zero é bem simples de entender e é o valor
que vai ser atribuído a uma variável quando você não inicializa ela com um valor, por exemplo se eu declarar uma `string` e
não colocar o seu valor ao printar esse valor em tela vou receber uma string em branco, se eu colocar um `int` ele vai me
dar 0. E de onde está vindo esse valor?
* No Go todo tipo de dado tem o seu valor zero que seria o seu valor inicial, então para `string` é uma string vazia, para
números sejam eles `int` ou `float` é o número 0, para `error` é nil que é o valor nulo no Go, então todo tipo de dado tem
o seu valor inicial e esse é o chamado valor zero.

## Booleano
* `true` ou `false`.
* Se você deixar sem declaração o valor zero do `bool` é `false`.

## Error
* O `error` em Go é um tipo, isso acontece porque o tratamento de erro dentro de Go é muito diferente do convencional e por
conta disso temos um tipo específico para erros em Go.
* O seu valor zero é `nil`.
* Como é que você cria um `error` em Go? Você tem que usar um pacote nativo no Go chamado `errors`