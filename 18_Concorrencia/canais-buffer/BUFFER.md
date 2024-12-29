## Canais com Buffer
* São canais um pouco diferentes desse canal que criamos na aula passada, aqui especificamos também uma capacidade para esse
canal.
* Perceba que nesse exemplo simples: criamos um canal, enviamos dados para esse canal, criamos uma variável para receber os
dados desse canal, exibimos em tela o valor dessa variável. Temos um deadlock e por que isso está acontecendo? É aquela questão
que a operação de receber e enviar dados são bloqueantes, então aqui eu vou enviar um valor para o canal o meu programa vai esperar
alguma outra linha receber um valor desse canal só que isso nunca vai acontecer porque ele vai ficar bloqueado aqui e nunca vai
chegar na linha em que a nossa variável está recebendo um valor ocasionando o deadlock.
* E é justamente por isso que a gente usa normalmente canais em funções separadas, então eu tenho uma função que envia um valor e
uma outra que recebe um valor, no exemplo da aula passada tínhamos a função escrever que estava enviando um valor e dentro da nossa
função main estávamos recebendo esse valor.
* Existe uma alternativa que poderíamos usar aqui para poder permanecer na mesma função que é criar um canal com buffer, é você
especificar uma capacidade para o canal.
* A diferença do canal com buffer e do canal normal (sem buffer) é que o canal com buffer ele só vai bloquear quando ele atingir a
sua capacidade máxima e o canal sem buffer não, ele vai sempre bloquear quando tiver uma operação de envio e recebimento.