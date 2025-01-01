## HTML
* A sigla quer dizer Hyper Text Markup Language, ou seja, linguagem de marcação de hipertexto. Isso quer dizer que você vai usar essa
linguagem como se fosse os blocos de construção para você fazer uma página na web, de um jeito resumido ele é a forma que você tem de
mostrar conteúdo dentro de uma página na internet.
* O html por si só é estático, ou seja, não consigo deixar ele dinâmico, por exemplo: apenas utilizando html não consigo fazer com que
a minha página faça uma saudação com o nome de usuário que está acessando ela, mas existem ferramentas/frameworks que nos auxiliam para
tornar o nosso html dinâmico e no Go existe um pacote que vai nos ajudar a fazer esse tipo de manipulação.
* Existe um pacote em Go chamado `"html/template"` e como o próprio nome sugere através dele a gente vai poder criar templates baseados
no nosso arquivo html e além da gente poder lidar com arquivos html no formato em que ele está a gente pode deixar ele dinâmico.