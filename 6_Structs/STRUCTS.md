## Structs
* O struct em Go é o que a linguagem tem de mais próximo de uma classe de uma linguagem de programação orientada a objetos.
* A definição mais formal é que ele é uma coleção de campos, cada campo tem um nome e um tipo, então imagine que você quer
salvar um usuário em uma variável, como você faria isso com os tipos normais que o Go te dá? Não teria como, imagine que o
usuário tem nome, idade, senha, email, você não pode salvar isso numa string, por exemplo, não tem como salvar todas essas
informações em uma string.
* Quando você tem vários campos, várias informações relativas a uma mesma coisa você usa o struct.
* Lembra quando comentamos sobre o valor zero, então o struct também tem o valor zero. Portanto, quando você cria um struct
e não passa nenhum dado o valor zero dele é bem simples, ele vai pegar todos os campos que estão dentro dele e vai jogar o
valor zero de cada um deles, assim de um campo string ele vai jogar uma string vazia, de um campo int ele vai jogar 0.
* Nada impede que tenhamos um struct dentro do outro, por exemplo posso criar um outro struct endereco e agora falar que a
minha struct usuario tem um endereco, desta forma estaríamos aninhando structs.