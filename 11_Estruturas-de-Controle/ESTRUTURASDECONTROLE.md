## Estruturas de Controle
* É muito simples você tem a avaliação de uma condição, se essa condição for verdadeira o seu programa faz uma coisa, se essa
condição for falsa o seu programa faz outra coisa.
* Não é necessário o uso de parênteses na checagem de condições pelo if, você pode estar utilizando os parênteses quando você
estiver fazendo várias operações que tenham uma precedência certa, então se você tiver três operações dentro do if e tiver que
ser executada em uma ordem certa você pode colocar elas entre parênteses.
* Algumas restrições: algumas linguagens permitem que você faça o if sem as chaves desde que você tenha uma condição de apenas
uma linha só de execução, no Go isso não funciona, da mesma forma você não pode deixar tudo na mesma linha. É preciso que seja
em linhas separadas e dentro de chaves.
* É uma estrutura muito similar ao que temos em outras linguagens, a única diferença que vamos ver aqui é o que chamamos de if
init que é quando você executa uma condição if e inicializa uma variável já nessa condição.
* O interessante desse if init é que se eu quiser acessar essa variável aqui fora não vai dar, vai nos informar que a variável
não está definida. Quando você cria uma variável usando if init você está limitando ela ao escopo do if, então fora desse escopo
ela não existe.
* Então, o if init é muito utilizado para isso às vezes você tem uma variável que vai ser utilizada só em uma condição e depois
vai ficar morta no programa você pode criar ela dessa maneira
* Recapitulando: o if vai avaliar uma condição, se ela for verdadeira vai executar o bloco de código que se encontra aqui, se
não for verdadeira ela vem para o else, e você pode aninhar condições com else if.