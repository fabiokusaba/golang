## Aplicação de Linha de Comando
* Uma rápida introdução uma aplicação de linha de comando ela seria uma aplicação comum que pode ser executada pela linha de comando.
* Até o presente momento nós rodamos os nossos arquivos com a instrução `go run <nome_do_arquivo>`, esse comando entra no pacote main
procura a função main e executa todas as instruções que estão ali dentro, a diferença para uma aplicação de linha de comando é que ela
pode receber inputs que a gente passa, então imagine que temos uma aplicação que tenha duas ações e poderíamos chamar essa aplicação da
seguinte forma: `go run main.go acao1 --parametro1 valor1 --parametro2 valor2`, e além disso eu poderia passar alguns parâmetros também.