# Star Wars API

Desenvolvido em Golang  (go version go1.17.6 darwin/amd64) e Elastic Search (version 7.17.0).

Utiliza a Star Wars API https://swapi.dev/api/ para complementar informações.

# Como Rodar

1. Ativar instância do Elastic Search
	$ elasticsearch
2. Buildar projeto (a partir da pasta do projeto)
	$ go build .
3. Ativar server Go (a partir da pasta do projeto) (default port 8083)
	$ go run . &
4. Executar chamadas via curl ou Postman
	$ curl http://localhost:8083/api/planets

# Serviços

 ### Adicionar Planeta

	api/planets POST

#### Parâmetros

name string
climate string
terrain string

#### Responses

201 OK 
{"status":true}
Filme salvo com sucesso

409 CONFLICT 
{"status":false, "message":"Planet already exists"}
Planeta não pode ser salvo, pois já existe

412 PRECONDITION FAILED 
{"status":false, "message":"All params must be set"}
Planeta não pode ser salvo, pois faltam parâmetros

500 INTERNAL SERVER ERROR
{"status":false, "message":"Error message"}

#### TODO (Task list)

- [ ] Validar atributos
- [ ] Validar se o planeta já existe
- [ ] Capturar total de aparições em filmes a partir da busca https://swapi.dev/api/planets/?search=Alderaan (caso o planeja não existir, salvar com 0 aparições)
- [ ] Criar Enum para climate e terrain


### Listar Planetas

	api/planets GET

#### Responses

200 OK 
[{"_id":"1", "name":"Tatooine", "climate": "arid", "terrain": "desert", "film_apparitions":"5"},{"_id":"2","name": "Alderaan", "climate": "temperate", "terrain": "grasslands, mountains", "film_apparitions":"2"}]

200 OK
[]
Nenhum planeta encontrado

#### TODO (Task list)

- [ ] Criar paginação na busca

### Buscar Planeta por nome

	api/planets?name=Tatooine GET

#### Responses

200 OK 
[{"_id":"1", "name":"Tatooine", "climate": "arid", "terrain": "desert", "film_apparitions":"5"}]

200 OK
[]
Nenhum planeta encontrado

#### TODO (Task list)

- [ ] Validar buscar com mais de uma palavra
- [ ] Implementar busca *full_text* no Elastic Search para sugestões aproximadas do nome

### Buscar Planeta por Id

	api/planets/1 GET

200 OK {"_id":"1", "name":"Tatooine", "climate": "arid", "terrain": "desert", "film_apparitions":"5"}

404 NOT FOUND

#### TODO (Task list)

- [ ] Validar Id

### Remover Planeta por Id

	api/planets/1 DELETE

#### Responses

200 OK 
{"status":true}
Planeta removido com sucesso

404 NOT FOUND 
{"status":false, "message":"Planet not found"}
Planeta não encontrado

#### TODO (Task list)

- [ ] Validar Id

### Respostas gerais de Erro

405 Method Not Allowed para outros métodos como HEAD, PUT e OPTIONS.
404 NOT FOUND para recursos não encontrado.

# Exemplo de Chamadas

	curl http://localhost:8083/api/planets

# Links de apoio

- https://httpstatuses.com/
- https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
- https://swapi.dev/api/planets/?search=Tatooine
- http://localhost:9200/ - Local Elastic Search
- https://swapi.dev/documentation