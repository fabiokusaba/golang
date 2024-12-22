## Maps
* É uma outra estrutura que você pode usar para guardar dados, tem algumas diferenças do struct principalmente porque ela é uma
estrutura de chave e valor mais rígida, a chave tem sempre o mesmo tipo, o valor também tem sempre o mesmo tipo e não é uma
estrutura mutável.
* Para criarmos um map é um pouco diferente, usamos a palavra `map`, dentro dos colchetes `[]` colocamos o tipo da nossa chave, na
sequência colocamos o tipo do nosso valor.
* Diferenças do map para o struct: não consigo fazer `usuario.nome` pois não tem esse campo para ele, se você quiser acessar a chave
nome fazemos da seguinte forma `usuario["nome"]`.
* Da mesma forma que poderíamos fazer um struct usuario aninhado podemos fazer com map também.
* Para deletar uma chave em um map basta usar uma função nativa no Go chamada delete, no primeiro parâmetro passamos o map que estamos
querendo deletar a chave e depois o nome da chave que você quer deletar.