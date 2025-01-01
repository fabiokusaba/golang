## JSON - Conceito e função Marshal
* Método para trafegar dados em uma aplicação.
* Formato que pode ser usado tanto para armazenar dados como para transportar dados de um lugar para o outro.
* Dentro do Go temos um pacote muito importante para lidarmos com JSON que é o: `encoding/json`.
* `json.Marshal()`: podemos usar para converter um Map para JSON ou um Struct para JSON.
* `json.Unmarshal()`: processo reverso, você transforma um JSON em um Struct ou em um Map dependendo da sua necessidade.
* O Go permite que logo após a declaração do tipo de dado do campo do seu struct você coloque: `json:"<nome_chave>"` que esse campo vai
representar dentro de um objeto JSON quando ele for convertido. Por convenção e boas práticas as chaves em JSON tem sempre letra minúscula.
* Quando usamos a função Marshal perceba que o resultado que obtemos é um slice de bytes ([]uint8), para conseguirmos ver esse resultado como
um JSON precisamos utilizar um outro pacote do Go que vai converter esse cara para uma visualização de JSON que conseguimos ver.
* Assim, vamos utilizar o `bytes.NewBuffer()` onde ele recebe um array de bytes e vai dar como saída um JSON.