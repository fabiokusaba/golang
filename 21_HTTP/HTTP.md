## Protocolo HTTP
* É um protocolo de comunicação usado para trafegar dados, de modo geral, e muitas vezes os dados que são trafegados são em formato JSON.
* Base da comunicação WEB.
* Funciona no modelo Cliente - Servidor, isso quer dizer que ele funciona num modelo onde o cliente faz uma requisição para um servidor,
esse servidor processa essa requisição e devolve uma resposta. Exemplo prático: imagine que você entre num e-commerce e você quer fazer o
cadastro nesse site, você entra no site, preenche os seus dados, você faz uma requisição para o servidor (site), o servidor vai processar
os dados que foram enviados e vai te dar uma resposta com base nesse processamento, então o servidor pode te dar uma resposta que o seu
cadastro foi concluído com sucesso ou que algum erro inesperado ocorreu.
* Request: termo em inglês para a requisição enviada pelo cliente.
* Response: termo em inglês para a resposta enviada pelo servidor após o processamento da requisição do cliente.
* Rotas: são uma maneira de você conseguir identificar o tipo de mensagem que está sendo enviada e também a partir disso identificar a que
tipo de processamento o servidor vai ter que fazer em cima dessa mensagem. Imagine o seguinte: ainda no exemplo do e-commerce você tem várias
coisas que podem ser feitas dentro do site como comprar um produto, avaliar um produto, criar uma conta, adicionar um produto ao carrinho de
compras, tem várias funcionalidades que você pode fazer e todas elas vão se comunicar com o servidor usando o mesmo protocolo HTTP, mas o que
muda é a requisição e a resposta porque para criar uma conta a requisição vai ser diferente do que a requisição que você faria para comprar um
produto, então a rota é usada para isso para identificar que ação deve ser feita e que tipo de mensagem você está mandando.
* As rotas tem dois componentes muito importantes que precisamos entender: um deles é o URI que seria o identificador do recurso então é uma forma
de você falar para o servidor do que você está falando, por exemplo: "/produtos", "/usuarios". E a segunda parte é o método, basicamente ele é
usado para que você fale o que você quer fazer com esse recurso que você acabou de identificar, por padrão temos os seguintes métodos: GET, POST,
PUT, DELETE.
* GET: buscar dados de um recurso.
* POST: cadastrar dados.
* PUT: atualizar dados.
* DELETE: remover dados.