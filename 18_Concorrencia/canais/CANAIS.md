## Canais
* Canais ou channels, em inglês, são a maneira mais utilizada que temos para a gente sincronizar as goroutines, para que a gente
possa fazer elas rodarem ao mesmo tempo de uma forma parecida com o que fizemos com WaitGroup só que de um jeito um pouci melhor.
* Esse nome deriva por ser uma canal de comunicação, então ele pode tanto enviar como receber dados e usamos dessa estratégia para
sincronizar as nossas goroutines.
* Para criarmos um canal usamos a função `make` passando `chan`e especificando o seu tipo, desta forma o canal criado só poderá
trafegar dados desse tipo especificado.
* O canal tem duas operações básicas muito importantes: enviar um dado e receber um dado. O que é interessante dessas operações é
que elas bloqueiam a execução do programa, então no nosso exemplo criamos um canal passamos ele para a nossa função e estamos
chamando essa função com a cláusula go, ou seja, o nosso programa não vai esperar essa função terminar de executar pra seguir, assim
ele chamou a função escrever e continuou.
* Aí ele chegou na linha onde o nosso canal está esperando receber um valor e vou salvar esse valor na variável mensagem, é justamente
nesse ponto que vai acontecer a sincronização com a goroutine da chamada da nossa função escrever porque quando falamos isso estamos
falando para ele obrigatoriamente esperar o canal receber um valor, e quando ele vai receber um valor? Quando alguém enviar um valor
para ele, no nosso caso isso ocorre dentro da função escrever.
* Então, o fluxo que está acontecendo é o seguinte: rodamos a função escrever com a cláusula go, ele vai começar a executar e já vai
embora, imagine que quando ele chegou na linha da nossa variável mensagem a função ainda não tinha executado, por exemplo, nesse caso
ele vai falar que ele precisa de um valor sendo enviado para o canal e vai ficar esperando esse valor chegar, somente quando esse valor
for enviado para o canal é que o programa vai passar dessa linha.
* Podemos pensar da seguinte forma: quando executamos a goroutine a gente não pede para ela esperar então ela começa a executar e vai
embora, quando chega no `<-canal` esse cara espera, é como se a goroutine começasse tudo sem esperar nada acontecer mas o canal precisa
esperar.
* A questão do nosso exemplo é que temos um for fazendo cinco execuções só que nesse caso específico o nosso programa não esperou nenhuma
delas porque ele esperou um valor chegar no canal e depois o go passa a não ligar se tem outras execuções para serem feitas, então a partir
do momento que o valor chegou para o canal ele vai continuar a execução e como depois do Println não tem mais nada para ser executado ele
termina o programa.
* Se eu quisesse fazer ele executar as cinco vezes, ou seja, fazer com que todas as mensagens fossem impressas na tela? Eu poderia fazer um
for desse cara.
* Deadlock -> é quando acontece o seguinte: você não tem mais nenhum lugar que está enviando dados para o seu canal só que o seu canal ainda
está esperando receber um dado. E isso é um problema porque o seu programa vai ficar ali eternamente esperando chegar um valor que nunca vai
chegar.
* Precisamos tomar muito cuidado com deadlock porque ele não é pego em compilação, o Go não consegue identificar um deadlock fazendo a 
compilação do código ele só pega em execução.
* Para tratarmos essa questão existe um jeito de verificarmos se o canal está aberto ou fechado, os canais possuem essa propriedade em Go, um
canal pode tanto estar aberto como fechado. Se o canal está aberto é sinal que ele ainda vai enviar dados e pode ainda receber dados, agora
se o canal estiver fechado é porque ele não vai nem enviar e nem receber dados.
* É importante você lembrar que são operações bloqueantes, então sempre que estou esperando receber um valor o meu programa está parado
esperando receber algo, por mais que seja uma função que está rodando de forma concorrente sempre que ele chega nesse momento de receber um
valor ele vai esperar alguém enviar algo, então você pode ter centenas de milhares de goroutines rodando ao mesmo tempo e usar os canais no
momento em que você precisa que uma delas retorne um valor, por exemplo.
* Uma outra coisa interessante é que além da gente usar essa propriedade aberto, booleano para verificar se o nosso canal está aberto ou
fechado, também temos uma outra maneira da gente conseguir fazer esse for sem precisar de um loop infinito.