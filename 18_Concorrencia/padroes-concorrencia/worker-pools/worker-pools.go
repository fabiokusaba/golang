package main

import "fmt"

func main() {
	// Para fazer a questão do worker pools vou precisar criar dois canais
	tarefas := make(chan int, 45)    // Sequência de números que precisamos calcular
	resultados := make(chan int, 45) // Armazenar os resultados a medida que forem sendo calculados

	// Chamando a função worker com uma goroutine, ele vai entrar na função, inicialmente o nosso canal de tarefas não tem nada, então
	// ao chegar no for ele vai esperar vir uma tarefa, desta forma conforme eu for mandando tarefas para esse cara ele vai realizando
	// o cálculo de fibonacci e mandando o valor para o canal de resultados
	// Eu posso chamar essa função worker quantas vezes eu quiser usando a questão da goroutine, então agora nesse caso eu vou ter dois
	// processos que vão estar puxando números dessa fila e fazendo as execuções
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	// Populando o canal de tarefas
	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	// Depois que o meu canal de tarefas já estiver preenchido eu vou fechar ele para que ele não possa receber mais nenhum tipo de dado
	close(tarefas)

	// Exibindo os resultados em tela
	for i := 0; i < 45; i++ {
		// A variável resultado vai receber o valor que vai ser enviado no canal de resultados, então ele vai ficar esperando um valor ser
		// enviado
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// Posso especificar na minha função o que esse canal irá fazer, ou seja, se ele só envia dados ou se ele só recebe dados
// Nesse caso o canal de tarefas vai ser o canal que somente recebe dados e o canal de resultados vai ser o canal que somente
// envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	// Vou iterar pelo meu canal de tarefas pegando todos os seus valores, calculando o número de fibonacci e jogando o resultado
	// no canal de resultados
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
