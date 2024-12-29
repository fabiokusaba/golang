## Padrões de Concorrência - Generator
* Basicamente esse cara ele vai ser uma função que vai encapsular uma chamada para uma goroutine e devolver um canal pra gente.
* A vantagem da nossa função escrever, que atende ao padrão generator, é que quando eu for chamar ela na função main eu não vou
precisar chamar ela com a cláusula go, posso chamar ela normalmente como uma função qualquer e ela quem vai encapsular uma chamada
para uma goroutine e vai me retornar um canal de comunicação.
* Então, utilizamos o generator para esconder essa complexidade, imagine que aqui poderia ter uma goroutine muito mais complexa ou
até mesmo uma série de goroutines e estou abstraindo tudo isso da função main, estou apenas chamando a função escrever de uma maneira
normal ela vai me retornar um canal e por esse canal que ela me retorna eu posso fazer a leitura dos dados.