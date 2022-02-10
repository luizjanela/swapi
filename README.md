# Star Wars API

Desenvolvido em Golang  (go version go1.17.6 darwin/amd64) e Elastic Search (version 7.17.0).

Utiliza a Star Wars API https://swapi.dev/api/ para complementar informações.

# Como Rodar

1. Ativar instância do Elastic Search (default port 9200)
	$ elasticsearch
2. Buildar projeto (a partir da pasta do projeto)
	$ go build .
3. Ativar server Go (a partir da pasta do projeto) (default port 8083)
	$ go run . &
4. Executar chamadas via curl ou Postman
	$ curl http://localhost:8083/api/planets
	
* Obs: Caso seja necessário customizar alguma porta ou endpoint, alterar no arquivo properties.go

# Entidade

A entidade Planeta possui 5 atributos:

- id string (gerado automaticamente)
- name string (inputado no serviço de criação)
- climate string (inputado no serviço de criação)
- terrain string (inputado no serviço de criação)
- apparitions int (capturado como int a partir do serviço SWAPI)

# Serviços

 ### Adicionar Planeta

	api/planets POST
	

#### Parâmetros

name string
climate string
terrain string

#### Responses

201 OK 
{"id":"_QRZ434BYh62MkvtPQvg", "name":"Tatooine", "climate": "arid", "terrain": "desert", "apparitions":5}
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

- [x] Validar atributos
- [x] Validar se o planeta já existe
- [x] Capturar total de aparições em filmes a partir da busca https://swapi.dev/api/planets/?search=Alderaan (caso o planeja não existir, salvar com 0 aparições)
- [ ] Implementar regra para tornar o serviço case-insensitive
- [ ] Criar Enum para climate e terrain
- [ ] Testes automatizados do serviço
- [x] Response geral para cenários de erro


### Listar Planetas

	api/planets GET

#### Responses

200 OK 
[{"id":"AgS4434BYh62MkvtxgxQ","name":"Tatooine","climate":"temperate","terrain":"gas giant","apparitions":5},{"id":"AwS5434BYh62MkvtbQzo","name":"Alderaan","climate":"temperate","terrain":"gas giant","apparitions":2}]

200 OK
[]
Nenhum planeta encontrado

500 INTERNAL SERVER ERROR
{"status":false, "message":"Error message"}

#### TODO (Task list)

- [ ] Criar paginação na busca
- [ ] Testes automatizados do serviço
- [x] Response geral para cenários de erro

### Buscar Planeta por nome

	api/planets?name={name} GET

#### Responses

200 OK 
[{"id":"AgS4434BYh62MkvtxgxQ","name":"Tatooine","climate":"temperate","terrain":"gas giant","apparitions":5}]

200 OK
[]
Nenhum planeta encontrado

500 INTERNAL SERVER ERROR
{"status":false, "message":"Error message"}

#### TODO (Task list)

- [x] Validar buscar com mais de uma palavra
- [ ] Implementar busca *full_text* no Elastic Search para sugestões aproximadas do nome
- [ ] Testes automatizados do serviço
- [x] Response geral para cenários de erro

### Buscar Planeta por Id

	api/planets/{id} GET

200 OK 
{"id":"AgS4434BYh62MkvtxgxQ","name":"Tatooine","climate":"temperate","terrain":"gas giant","apparitions":5}

404 NOT FOUND
{"status":false,"message":"Planet not found"}

500 INTERNAL SERVER ERROR
{"status":false, "message":"Error message"}


#### TODO (Task list)

- [x] Validar Id
- [ ] Testes automatizados do serviço
- [x] Response geral para cenários de erro

### Remover Planeta por Id

	api/planets/{id} DELETE

#### Responses

200 OK 
{"status":true, "message":"Planet deleted"}
Planeta removido com sucesso

404 NOT FOUND 
{"status":false, "message":"Planet not found"}
Planeta não encontrado

#### TODO (Task list)

- [x] Validar Id
- [ ] Testes automatizados do serviço
- [x] Response geral para cenários de erro

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