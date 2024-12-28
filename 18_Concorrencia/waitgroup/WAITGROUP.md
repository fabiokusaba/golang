## Waitgroup
* Maneiras de sincronizar goroutines: imagine que tenho um programa que tem três funções que podem ser executadas ao mesmo tempo
só que para eu terminar esse programa eu preciso que as três funções tenham terminado, então eu começo as três funções ao mesmo
tempo, não tenho a dependência de uma esperar a outra, mas no final do programa eu preciso esperar as três.
* Existem maneiras da gente fazer essa sincronização entre as goroutines e uma das técnicas é chamada de `waitgroup`, uma outra
técnica seria os `channels`, canais, que seria a solução mais adequada para a grande maioria dos casos.
* Imagine esse seguinte cenário: vou chamar a função `escrever` duas vezes e depois disso eu quero que o meu programa termine, mas
eu não quero que essas duas funções dependam uma da outra para executar, ou seja, eu não quero que a segunda chamada dependa da
primeira chamada terminar para executar, então eu quero fazer isso de um jeito um pouco mais dinâmico/rápido, nesse caso podemos
usar as goroutines mas eu quero sincronizá-las, ou seja, quero que as duas tenham terminado antes do meu programa terminar de executar.
* Então, para isso podemos criar um `waitgroup` e como esse cara funciona? Ele funciona da seguinte forma: logo após a sua declaração
eu uso o seguinte comando `Add` colocando a quantidade de goroutines que vão estar rodando ao mesmo tempo, em tradução ele seria um
grupo de espera então estou passando para ele que tem duas goroutines que fazem parte desse grupo de espera.
* Posso aplicar uma estratégia um pouco diferente para rodar essas duas chamadas de funções em paralelo, então vou chamar uma função
anônima com a cláusula `go` para que ela vire uma goroutine e aí vou chamar a função que quero executar dentro dela.
* Quando essa função terminar vou chamar o `waitGroup.Done()`, finalizada as chamadas das funções eu vou chamar um `waitGroup.Wait()`.