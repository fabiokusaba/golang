## Introdução às Goroutines
* Diferença entre concorrência e paralelismo: o paralelismo acontece quando você tem duas ou mais tarefas que estão sendo
executadas exatamente ao mesmo tempo isso só é possível se você tiver um processador com mais de um núcleo que aí ele
consegue executar uma tarefa em cada núcleo, ele distribui todas essas tarefas entre esses núcleos e aí sim elas vão estar
sendo executadas ao mesmo tempo.
* Agora tarefas que executam de forma concorrente elas não necessariamente estão executando ao mesmo tempo, elas podem estar
também se você tiver, mesmo caso, mais de um núcleo no processador que nesse caso elas seriam divididas entre os núcleos, mas
não necessariamente, se você tiver um processador com apenas um núcleo também é possível aplicar concorrência nele.
* A diferença maior seria que você teria, por exemplo, duas tarefas que estão rodando e uma tarefa não esperaria a outra acabar
pra rodar também, então elas ficariam, de certa forma, revezando tempo.
* Uma goroutine é uma função/método que é invocado colocando a palavra `go` na frente da função que você está chamando/executando.
* Quando você coloca a palavra `go` na frente de uma chamada de função ou de um método o que você está falando para o seu código?
Você está falando o seguinte: executa essa função que eu assimilei com `go`, mas não espera ela terminar pra seguir o programa.
* Você pode estar se perguntando que: "eu poderia colocar goroutines o tempo inteiro no meu programa e aí ele ficaria mil vezes
mais rápido porque ele não vai esperar que uma determinada ação termine pra começar outra". Porém, ao fazermos isso percebemos que
o programa simplesmente finaliza e o por quê disso? Porque quando você coloca `go` na frente da chamada de uma função ou método você
está falando para o seu programa não esperar terminar a execução e passar para a linha seguinte, só que na linha seguinte você falou
a mesma coisa para o seu programa, então como não tem mais nada para ele executar ele vai chegar ao final da função main e vai 
finalizar sua execução sem exibir nada na tela.