# quake-log-parser

## Solução
* Criei um CLI para facilitar o uso da ferramenta.
* Fiz o parse linha a linha, pensando que o log poderia crescer e não seria legal colocar ele todo em buffer.
* Sobre a API, está muito simples, servindo apenas o relatório como solicitado.
* Testei todo o necessário para o funcionamento do software, usando TDD. Não fiz testes únitarios para os handlers, mas compensei disponibilizando [a collection que usei para testar no postman.](https://github.com/zeucxb/quake-log-parser/blob/master/quake-log-parser.postman_collection.json)
* Não utilizei banco de dados, para simplificar a execução do teste para análise.

## Setup
Para testar, basta compilar o programa e usar o CLI:

```bash
// você pode usar ./quake-log-parser h (ou helper) para pedir ajuda
// também pode usar go run main.go no lugar de ./quake-log-parser se resolver não compilar o CLI

$ go build ./
$ ./quake-log-parser p // ou parse
$ ./quake-log-parser s // ou serve
```

O método serve (s) aceita uma porta por parametro, seu default é [localhost:3000](http://localhost:3000)

## Relatório (endpoints)
Método| Rota                | Exemplo                                              
----- | ------------------- | -----------------------------------------------------
(GET) | /                   | [/](http://localhost:3000)                           
(GET) | /report             | [/report](http://localhost:3000/report)              
(GET) | /report/:gameName   | [/report/game_1](http://localhost:3000/report/game_1)
(GET) | /report/:gameNumber | [/report/1](http://localhost:3000/report/1)          

Você também pode acessar [a collection que usei para testar no postman.](https://github.com/zeucxb/quake-log-parser/blob/master/quake-log-parser.postman_collection.json)
