## Select
* É uma estrutura muito parecida com switch, que já vimos para tomada de decisão, só que ela é focada para uso com concorrência.
* Perceba que no exemplo que criamos não está dando erro e está funcionando corretamente, porém estamos perdendo tempo porque a
primeira goroutine envia um valor para o canal a cada meio segundo isso quer dizer que essa variável que está recebendo o valor
do canal teoricamente poderia receber um valor a cada meio segundo só que isso não está acontecendo porque ele está chegando na
segunda variável que recebe valor de outro canal então vou estar esperando ele receber um valor para assim dar continuidado ao
loop.
* E como podemos fazer para parar de perder esse tempo? A gente usa o select.
* Em resumo o select é usado para concorrência em um cenário que você tem uma função que demora muito mais tempo para terminar do
que a outra ele pode ser uma alternativa muito boa.