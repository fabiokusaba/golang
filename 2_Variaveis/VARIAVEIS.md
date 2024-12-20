## Variáveis
* Em Go existem duas maneiras de você declarar uma variável: a primeira maneira é você deixando o tipo dela explícito e a segunda é você
deixando o tipo dela implícito.
* O Go é uma linguagem fortemente tipada, então todas as variáveis tem obrigatoriamente um tipo.
* O Go não deixa você declarar variáveis e não usá-las isso reflete que o seu código não irá compilar caso você tenha uma variável declarada
e não utilizada, essa mesma regra é aplicada para imports, ou seja, se você importar um pacote/biblioteca e não utilizar ele também não vai
deixar o seu código rodar.
* Na declaração implícita o Go determina o tipo da variável com base no seu valor. Isso também é chamado de inferência de tipo, ou seja, você
está inferindo o tipo de uma variável com base no seu valor.
* A diferença entre uma constante e uma variável é que na constante você não pode mudar o valor dela.
* Uma coisa que é muito interessante no Go é a forma como conseguimos inverter o valor de duas variáveis em outras linguagens de programação
quando nos deparamos com esse desafio geralmente criamos uma variável temporária para armazenar o valor de um deles para assim conseguirmos
realizar a troca, porém em Go isso é feito de forma muito simplificada veja no exemplo da `variável5` e `variável6`.