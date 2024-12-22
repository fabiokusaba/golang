## Ponteiros
* Você pode pensar no ponteiro como sendo uma variável que vai salvar dentro dela não necessariamente um valor, mas o
endereço de memória de alguma coisa.
* Quando você atribui um valor para uma variável, por padrão, esse valor é uma cópia.
* Qual a diferença de você atribuir valores assim e atribuir valores usando ponteiros? O ponteiro é uma referência de
memória, então quando você cria um ponteiro e você coloca, por exemplo, a variavel2 que a gente criou nele, na verdade
você não vai estar jogando o valor da variável dentro dele você vai estar jogando o endereço de memória onde essa variável
está salva, porque quando você cria uma variável em seu programa o seu computador vai lá e cria um endereço de memória dentro da memória dele e coloca a variável nesse lugar como se fosse uma gaveta.
* Imagine que a memória do seu computador tem várias gavetas e você fala para o computador criar uma variável, ele vai pegar e vai colocar em uma dessas gavetas.
* Quando você cria um ponteiro e você passa uma variável para esse ponteiro você não está passando o que está dentro da gaveta você está passando como se fosse o endereço da gaveta.
* Desreferenciação é o processo de desfazer a referência, então você vai estar lendo o valor que está dentro desse endereço
de memória.
* Por quê ponteiros são relevantes? Imagine que eu vou mudar o valor da variavel3, perceba que não houve mudança no ponteiro, ou seja, não houve mudança no seu endereço de memória. Se olharmos para os valores percebemos que ambos estão com
o mesmo valor, pois o valor que está armazenado nesse endereço de memória é diferente mas o nosso ponteiro ainda está apontando para ele, então se antes tinha um valor 100 e agora tem um valor 150 o ponteiro vai conseguir pegar porque ele está apontando para o mesmo endereço.