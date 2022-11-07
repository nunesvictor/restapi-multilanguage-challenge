# Go vs Python vs Rust REST API challenge

O objetivo desse projeto é fazer um mesmo sistema em [Go](https://go.dev/), [Python](https://www.python.org/) e em [Rust](https://www.rust-lang.org/).

A partir do desenvolvimento de uma mesma aplicação utilizando as ferramentas acima citadas a intenção é de analizar os prós e os contras de cada uma, sendo levado em consideração a praticidade no desenvolvimento, dificuldade geral de implementação e, claro, desempenho.

A motivação vem da tentativa de encontrar alternativas de alto nível para o stack Django + Django REST Framework uma vez que essa combinação apesar de poderosa em muito tem se mostrado com desempenho abaixo do esperado para requisições complexas ou fortemente carregadas de relacionamentos.

___

## Das ferramentas de desenvolvimento

Para o desenvolvimento do propõe-se as seguintes ferramentas:

### GO

Aplicação Go 1.19.3 (darwin/arm64) construída com [Gin](https://gin-gonic.com/)  Web Framework e ORM [Gorm](https://gorm.io/).

### Python

Aplicação Python 3.11 (darwin/arm64) construída com [FastAPI](https://fastapi.tiangolo.com/) Web Framework e ORM [SQLAchemy](https://www.sqlalchemy.org/).

### Rust

Aplicação Rust 1.65.0 (darwin/arm64) construída com [Actix](https://actix.rs/)  Web Framework e ORM [Diesel](https://diesel.rs/).

### Database

Para todas as aplicações será utilizada a base de dados [PostgreSQL](https://www.postgresql.org/) 15 (darwin/arm64).

___

## Do projeto

### Estrutura de diretórios

Na raiz desse projeto serão criadas 3 (três pastas): `go`, `python` e `rust`, pastas estas que representarão a raiz de cada um dos respectivos projetos web.

### Informações gerais

O sistema que se propõe é um sistema de controle de biblioteca. Pretende-se criar uma API completa para o sistema que manipule a inserção, alteração, listagem e exclusão de livros bem como todos os seus relacionamentos.

O sistema também deve exigir que o acesso à API seja autenticado (via `Authorization` header, **NÃO SERÁ PERMITIDO** autenticação via `Basic Autentication`).

Será exigido a existência de apenas um usuário administrador que terá permissão para realizar as operações básicas (*CRUD*) em todos os *resources* da API.

Segue abaixo a lista das tabelas que deverão ser criadas e manipuladas pela aplicação:

| Livro                  |
|------------------------|
| + id: int              |
| + titulo: string       |
| + editora: Editora     |
| + autores: Autor[]     |
| + generos: Genero[]    |
| + edicao: int          |
| + ano: int             |
| + paginas: int         |
| + ISBN: string         |
| + created_at: DateTime |
| + updated_at: DateTime |

| Editora                |
|------------------------|
| + id: int              |
| + nome: string         |
| + localidade: string   |
| + created_at: DateTime |
| + updated_at: DateTime |

| Autor                  |
|------------------------|
| + id: int              |
| + nome: string         |
| + sobrenome: string    |
| + created_at: DateTime |
| + updated_at: DateTime |

| Genero                 |
|------------------------|
| + id: int              |
| + descricao: string    |
| + created_at: DateTime |
| + updated_at: DateTime |

## Das funcionalidades adicionais

Deverá haver a possibilidade de buscar Livros, Autores e Editoras por uma query digitada pelo usuário, devendo o sistema pesquisar pela coluna específica no banco de dados.

O sistema deverá também fornecer as seguintes consultas:

1. Livros por autor
2. Livros por editora
3. Livros por gênero

## Da documentação

Todas os projetos deverão fornecer documentação padrão `OpenAPI` (swagger).

## Das análises de desempenho e produtividade

Será considerado o tempo total para o desenvolvimento para cada uma das aplicações. Pensa-se ser uma boa métrica pois o desenvolvedor desconhece quaisquer das ferramentas usadas, logo, o próprio tempo necessário para percorrer a curva de aprendizagem da ferramenta mostra-se válido como unidade de medida de produtividade (pelo menos nesse primeiro momento).

Ainda sobre produtividade, a complexidade na implentação da documentação `OpenAPI` será levada em consideração.

Para desempenho, será considerado apenas o tempo de execução de cada uma das operações de manipulação do *resource*. Serão feitos vários testes e a métrica se baseará na média das respostas de cada bateria de testes.