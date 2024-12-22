## Funções
* Funções também são um tipo em Go, um tipo um pouco diferente dos outros.
* Em resumo uma função é como se fosse uma receita de bolo, ou seja, uma série de passos que vão ser seguidos pelo seu
programa.
* As funções em Go podem ter parâmetros e retornos, então parâmetros são dados que você coloca na função e retorno é o que
a função devolve pra você.
* Para declararmos uma função em Go utilizamos a palavra reservada `func` seguido do nome da função, entre parênteses 
colocamos os parâmetros que a nossa função irá receber, na sequência colocamos ou não o seu retorno, abrimos as chaves e 
aqui dentro, no corpo da função, escrevemos a lógica que queremos que ela execute.
* As funções são tipos, então eu posso tanto igualar uma variável a uma função e eu também posso fazer com que o parâmetro
de uma função seja outra função, que o retorno de uma função seja outra função, dá para aninharmos essas coisas.
* O tipo da função é composto pela palavra chave, sempre, pelos parâmetros e pelo retorno.
* As funções em Go podem ter mais de um retorno.
* Quando temos múltiplos retornos em uma função e não precisamos de todos eles podemos descartar esse valor utilizando `_`,
underline/underscore, e aqui a ordem é muito importante e pode alterar o resultado desejado, pois a ordem vai seguir o
retorno da função.