
#  Star  Wars  API

  

Desenvolvido  em  Golang  (go  version  go1.17.6  darwin/amd64)  e  Elastic  Search  (version  7.17.0).

  

Utiliza  a  Star  Wars  API  https://swapi.dev/api/  para  complementar  informações.

  

#  Arquitetura

  

O  projeto  se  baseia  no  padrão  MVC  (sem  View)  e  no  padrão  REST.  Foi  criada  uma  nova  camada  para  abstrair  a  comunicação  com  os  provedores  de  dados  externos  chamada  Library.  A  camada  library  (`elklibrary.go`  e  `swlibrary.go`)  entregam  dados  já  no  modelo  definido,  fazendo  com  que  o  Controller  apenas  tenha  que  se  preocupar  com  a  camada  de  comunicação  HTTP.

  

A  Library  `elklibrary.go`  é  reponsável  por  comunicar  com  o  Elastic  Search.

A  Library  `swlibrary.go`  é  reponsável  por  comunicar  com  o  serviço  SWAPI  (unicamente  usado  para  preencher  o  dado  de  apparitions).

  

Apenas  temos  o  controller  `planetcontroller.go`  pois  planet  é  o  único  recurso  tratado  pelo  projeto.  Da  mesma  forma,  o  `planetmodel.go`  define  o  model  do  Planeta.

  

Para  o  tratamento  de  algumas  respostas,  foi  criada  também  a  struct/model  `responsemodel.go`  para  definição  do  padrão  de  resposta  com  o  atributo  status.

  

Foi  criado  também  o  arquivo  `properties.go`  para  que  configurações  gerais  e  strings  padrão  fossem  inicializadas  em  um  ponto  central  através  de  constantes.

  

O  arquivo  `server.go`  é  onde  está  presente  o  método  main,  sendo  ele  executado  o  método  executado  ao  rodar  o  comando  `go  run  .  &`.

  

Após  o  projeto  iniciar,  podem  ser  feitas  chamadas  HTTP  para  que  acesso  aos  recursos.

O projeto se utiliza do package `github.com/gorilla/mux` para facilitação da criação de rotas e handlers.
  

#  Como  Rodar

  

Ativar  instância  local  do  Elastic  Search  (port  9200)

  `$  elasticsearch`

Buildar  projeto  (a  partir  da  pasta  do  projeto)

 `$  go  build  .`

Ativar  server  Go  (a  partir  da  pasta  do  projeto)  (port  8083)

  `$  go  run  .  &`

Executar  chamadas  via  curl  ou  Postman

 ` $  curl  http://localhost:8083/api/planets -v`

Obs:  Caso  seja  necessário  customizar  alguma  porta  ou  endpoint,  alterar  no  arquivo  properties.go

  

#  Planeta

  

A  entidade  Planeta  possui  5  atributos:

 

-  `id  string`  (gerado  automaticamente)

-  `name  string`  (inputado  no  serviço  de  criação)

-  `climate  string`  (inputado  no  serviço  de  criação)

-  `terrain  string`  (inputado  no  serviço  de  criação)

-  `apparitions  int`  (capturado  como  int  a  partir  do  serviço  SWAPI)

  

#  Serviços

  

  ###  Adicionar  Planeta

  

`  api/planets  POST`

  

####  Parâmetros

  

 - `name  string` nome do Planeta
 - `climate  string` clima do Planeta
 - `terrain  string` terreno do Planeta

  

####  Responses

  
| HTTP Status | Body | Descrição |
|--|--|--|
| 201  OK | {"id":"_QRZ434BYh62MkvtPQvg",  "name":"Tatooine",  "climate":  "arid",  "terrain":  "desert",  "apparitions":5} | Filme  salvo  com  sucesso|
|409  CONFLICT |   {"status":false,  "message":"Planet  already  exists"} | Planeta  não  pode  ser  salvo,  pois  já  existe |
| 412  PRECONDITION  FAILED | {"status":false,  "message":"All  params  must  be  set"} | Planeta  não  pode  ser  salvo,  pois  faltam  parâmetros |
| 500  INTERNAL  SERVER  ERROR | {"status":false,  "message":"Error  message"} | General error |

  

####  TODO  (Task  list)

  

-  [x]  Validar  atributos

-  [x]  Validar  se  o  planeta  já  existe

-  [x]  Capturar  total  de  aparições  em  filmes  a  partir  da  busca  https://swapi.dev/api/planets/?search=Alderaan  (caso  o  planeja  não  existir,  salvar  com  0  aparições)

-  [  ]  Implementar  regra  para  tornar  o  serviço  case-insensitive

-  [  ]  Criar  Enum  para  climate  e  terrain

-  [  ]  Testes  automatizados  do  serviço

-  [x]  Response  geral  para  cenários  de  erro

  

  

###  Listar  Planetas

  

 `api/planets  GET`

  

####  Responses

  


| HTTP Status | Body | Descrição |
|--|--|--|
| 200  OK  | [{"id":"AgS4434BYh62MkvtxgxQ","name":"Tatooine","climate":"temperate","terrain":"gas  giant","apparitions":5},{"id":"AwS5434BYh62MkvtbQzo","name":"Alderaan","climate":"temperate","terrain":"gas  giant","apparitions":2}] | Listagem dos Planetas  |
| 200  OK | [] | Nenhum  planeta  encontrado |
| 500  INTERNAL  SERVER  ERROR | {"status":false,  "message":"Error  message"} | General error | 

  

####  TODO  (Task  list)

  

-  [  ]  Criar  paginação  na  busca

-  [  ]  Testes  automatizados  do  serviço

-  [x]  Response  geral  para  cenários  de  erro

  

###  Buscar  Planeta  por  nome

  

  `api/planets?name={name}  GET`



####  Parâmetros

  

 - `name  string` Nome do planeta a ser buscado
  

####  Responses

  


| HTTP Status | Body | Descrição |
|--|--|--|
|200  OK  | {"id":"AgS4434BYh62MkvtxgxQ","name":"Tatooine","climate":"temperate","terrain":"gas  giant","apparitions":5} | Planeta encontrado  |
|404  NOT FOUND | {"status":false,  "message":"Planet not found"} | Nenhum  planeta  encontrado |
|500  INTERNAL  SERVER  ERROR | {"status":false,  "message":"Error  message"} | General error |

  

####  TODO  (Task  list)

  

-  [x]  Validar  buscar  com  mais  de  uma  palavra

-  [  ]  Implementar  busca  *full_text*  no  Elastic  Search  para  sugestões  aproximadas  do  nome

-  [  ]  Testes  automatizados  do  serviço

-  [x]  Response  geral  para  cenários  de  erro

  

###  Buscar  Planeta  por  Id

  

  `api/planets/{id}  GET`



####  Parâmetros

  

 - `id  string` Id do planeta a ser buscado

#### Responses



| HTTP Status | Body | Descrição |
|--|--|--|
| 200  OK  | {"id":"AgS4434BYh62MkvtxgxQ","name":"Tatooine","climate":"temperate","terrain":"gas  giant","apparitions":5} | Planeta encontrado | 
| 404  NOT  FOUND | {"status":false,"message":"Planet not found"} |  Planeta não encontrado | 
|500  INTERNAL  SERVER  ERROR|{"status":false,  "message":"Error  message"}|General error|

  

  

####  TODO  (Task  list)

  

-  [x]  Validar  Id

-  [  ]  Testes  automatizados  do  serviço

-  [x]  Response  geral  para  cenários  de  erro

  

###  Remover  Planeta  por  Id

  

  `api/planets/{id}  DELETE`



####  Parâmetros

  

 - `id  string` Nome do planeta a ser deletado
  

####  Responses

  


| HTTP Status | Body | Descrição |
|--|--|--|

|200  OK  |{"status":true,  "message":"Planet  deleted"}|Planeta  removido  com  sucesso|
  
| 404  NOT  FOUND | {"status":false,"message":"Planet not found"} |  Planeta não encontrado | 

|500  INTERNAL  SERVER  ERROR|{"status":false,  "message":"Error  message"}|General error|

  

####  TODO  (Task  list)

  

-  [x]  Validar  Id

-  [  ]  Testes  automatizados  do  serviço

-  [x]  Response  geral  para  cenários  de  erro

  

###  Respostas  gerais  de  Erro

  
| HTTP Status | Descrição |
|--|--|

|405  METHOD NOT ALLOWED |Método não permitido|
  
| 404  NOT  FOUND | Recurso não encontrado | 


  

#  Exemplo  de  Chamadas

  
Listagem de Planetas
`curl -X GET http://localhost:8083/api/planets`

Criação de Planeta
`curl -X POST http://localhost:8083/api/planets --data "name=Alderaan&climate=temperate&terrain=gas giant" -v`
  

#  Links  de  apoio

  

-  https://httpstatuses.com/
- https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
-  https://swapi.dev/documentation
- https://pkg.go.dev/