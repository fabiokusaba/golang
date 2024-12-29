## Padrões de Concorrência - Worker Pools
* Basicamente é como se você tivesse uma fila de tarefas para serem executadas e você tem vários workers/processos pegando itens
dessa fila de uma maneira independente.
* Imagine a fila como sendo vários números de fibonacci pra gente calcular, então teria mais de um processo rodando ao mesmo tempo
que pegaria um número dessa fila e iria fazendo o cálculo do número de fibonacci.