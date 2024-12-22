## Arrays e Slices
* O array no Go como em qualquer outra linguagem de programação é nada mais do que uma lista de valores, então você tem
uma mesma variável que vai guardar uma série de valores.
* Existem duas formas de se declarar arrays em Go: uma delas é usando `var` seguido do seu nome, tamanho e tipo. Uma outra
forma é usando a inferência de tipos e já passar o seu valor.
* Lembrando que o Go não vai deixar você criar uma coisa de qualquer tipo e também quando você cria um array todos os itens
desse array tem que ser obrigatoriamente do mesmo tipo.
* Quando declaramos um array e não passamos valores para dentro dele ao printarmos na tela veremos o seu valor zero.
* Os arrays sempre iniciam da posição 0, portanto um array de 5 posições vai ser: 0 - 1 - 2 - 3 - 4, igual vemos em outras
linguagens.
* Obrigatoriamente precisamos especificar o tamanho do array, caso contrário veremos que ele irá virar um slice.
* O array é bem inflexível nesse sentido, pois uma vez declarado o seu tamanho não podemos adicionar elementos a mais e
também além de respeitarmos o seu tamanho fixo precisamos respeitar o tipo de dados que foi declarado.
* O array não é muito utilizado no Go, normalmente nas aplicações que vamos desenvolver utilizamos mais o slice que é uma
estrutura feita baseada no array só que ela tem algumas flexibilidades a mais principalmente na questão do tamanho.
* Para a gente criar um slice é bem parecido do que faríamos para criar um array, a diferença maior está na questão dos
colchetes, aqui não vou especificar um tamanho, então ele não tem um tamanho fixo e o seu tamanho pode mudar de acordo com
a minha necessidade.
* Os slices ainda possuam a limitação do tipo, ou seja, se eu criar um slice de inteiros ele só pode ter números inteiros não podendo ter nenhum dado de outro tipo dentro dele.
* É importante entender que o slice não é um array, ele parece muito com um array, na verdade ele aponta para um array, ele
tem um componente dele que é como se fosse um ponteiro que aponta para um array que ele referencia, daí o seu nome slice que no inglês quer dizer fatia, é como se fosse inicialmente uma fatia de um array.
* No slice temos uma função chamada append e o que ela faz é basicamente adicionar um item no slice e retornar um slice novo com esse item adicionado. O primeiro parâmetro passamos o slice e no segundo parâmetro o valor que queremos adicionar
então essa função vai pegar o slice que passamos vai adicionar o item e vai retornar um slice novo com todas as informações
que já existiam nesse slice mais esse item, podemos jogar na mesma variável e desta forma é como se estivéssemos alterando
propriamente o nosso slice.
* O slice é uma fatia, então ele é um pedaço de um array portanto conseguimos pegar um intervalo do nosso array e jogar em
um slice, o array vai ficar inalterado e vou pegar um pedaço dele e salvar dentro do meu slice. Importante ressaltar que
essa sintaxe o primeiro índice é incluso e segundo é exclusivo, ou seja, ele não vai pegar o item que esteja nesse índice
que foi passado pegando o valor anterior a ele.
* Então, nesse caso criamos um slice (fatia) a partir de um array que já existia, esse cara aqui é como se fosse um ponteiro apontando para essas posições do array.

## Arrays Internos
* A função make vai alocar um espaço na memória para uma determinada coisa do nosso programa, então pode ser para um slice como vai ser o nosso caso, você pode criar
outras estruturas no Go também usando a função make como alternativa.
* A função make recebe três parâmetros: o primeiro deles vai ser o tipo do slice que queremos criar, o segundo parâmetro vai ter o tamanho do nosso slice que é a quantidade
de itens que ele tem, e o terceiro parâmetro é a capacidade que é a quantidade máxima de itens que ele pode ter.
* Quando você cria um slice ele está sempre apontando para um array daí vem o nome slice que é uma fatia de um array, o que acontece é que essa função make ela vai criar um
array de 15 posições e ela vai retornar pra gente um slice que vai pegar as 10 primeiras posições desse array, então quando você usa um slice, quando você cria ele sem fazer
uma referência direta a um array ele vai referenciar o que a gente chama de array interno.
* Nesse caso, a função make criou um array pra gente de 15 posições, por isso capacidade, e retornou pra gente um slice de 10 posições que vai referenciar as 10 primeiras posições
desse array.
* O Go quando ele vê que o seu slice vai estourar o tamanho ele cria um outro array pra você referenciar e dobra o tamanho dele e é por isso que o slice nunca tem um tamanho limite
ele tem a capacidade máxima dele inicialmente, só que sempre que você passar essa capacidade ele vai te dar mais, então ele nunca vai deixar que você fique sem espaço.
* Basicamente quando você não passa o último parâmetro aqui no make ele vai definir que a capacidade do slice é igual ao tamanho dele.