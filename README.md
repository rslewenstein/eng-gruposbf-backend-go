# eng-gruposbf-backend-go
Teste para Engenheiro de Software do Grupo SBF.
#### 1 - Baixar as dependências:
- Executar o comando: 
```go mod tidy```

- versão do Go: 1.19

#### 2 - criar o arquivo .env, dentro da diretório api. (Já existe um arquivo de EXEMPLO_ENV, basta renomea-lo para .env e preencher os dados):
- DB_USER= (usuario do BD).
- DB_PASSWORD= (senha do BD).
- DB_NAME= (nome do BD).
- API_PORT= (Localmente eu coloco API_PORT=5000, mas se nenhuma porta for setada, automaticamente será atribuída a porta 9000).
- API_KEY= (Chave da API de conversão de moedas).

#### 3 - Criar o banco de dados, a tabela e os inserts:
- No diretório ```api/misc/scripts```, contém os scripts para execução.
- Executar primeiro o script ```CreateDatabaseTables.sql``` e em seguida, execute o script ```inserts.sql```

- https://exchangeratesapi.io/documentation/ (expirou)

- https://openexchangerates.org/ (https://docs.openexchangerates.org/reference/api-introduction)


- http://localhost:5000/swagger/index.html
- https://github.com/swaggo/swag
- https://github.com/swaggo/http-swagger/tree/master/example/gorilla