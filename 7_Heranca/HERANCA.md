## "Herança"
* Vamos trabalhar com o mais perto que Go chega de herança convencional em uma linguagem orientada a objetos, visto que o
Go não é uma linguagem orientada a objetos propriamente dita.
* Criamos duas structs: uma pessoa e um estudante, fazendo um paralelo com o mundo real todo estudante é uma pessoa, então
se você tiver um struct com todas as características que englobam uma pessoa você poderia também passar as características
para o estudante, sendo que o estudante teria algumas características a mais.
* Desta forma, para evitarmos ter que repetir as mesmas características de pessoa em estudante basta passarmos apenas o seu
nome sem passar um tipo pra ele, basicamente essa é a forma de fazermos herança em Go. Assim, o que estamos fazendo é pegar
todos os campos que estão na struct pessoa e jogando dentro da struct estudante, e você pode acessar esses campos da mesma
forma que você faria com a pessoa.
* Esse é o único caso que você não precisa especificar o tipo aqui, na verdade isso aqui já é o tipo então você não precisa
fazer `pessoa pessoa`.