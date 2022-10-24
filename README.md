# eng-gruposbf-backend-go
Teste para Engenheiro de Software do Grupo SBF.
- Candidato: Rafael Soares Lewenstein
#### 1 - Dependências:
- versão do Go: 1.19
- Utilizei o MySql

- Executar o comando: 
```go mod tidy```

#### 2 - criar o arquivo .env, dentro da diretório api. (Já existe um arquivo de EXEMPLO_ENV, basta renomea-lo para .env e preencher os dados):
- DB_USER= (usuario do BD).
- DB_PASSWORD= (senha do BD).
- DB_NAME= (nome do BD).
- API_PORT= (Localmente eu coloco API_PORT=5000, mas se nenhuma porta for setada, automaticamente será atribuída a porta 9000).
- API_KEY= (Chave da API de conversão de moedas). Será necessário gerar a chave para utilizar a API de terceiros (https://exchangeratesapi.io/documentation/)

#### 3 - Criar o banco de dados, a tabela e os inserts:
- (tentei utilizar migrations para automatizar o processo, mas não funcionou bem).
- No diretório ```api/misc/scripts```, contém os scripts para execução.
- Executar primeiro o script ```CreateDatabaseTables.sql``` e em seguida, execute o script ```inserts.sql```

#### 4 - Endpoints:

- Para converter a partir de uma moeda e um valor informado (```GET```):
- ```"/api/convert/{symbol}&{amount}" ```
- EX: ``` localhost:5000/api/convert/BRL&500 ```

- Para cadastrar uma moeda e um país (```POST```):
- ``` "/api/convert/currency" ```
- EX: 
```
 "localhost:5000/api/convert/currency"
{
    "symbol": "GBP",
    "country": "Reino Unido"
}
```

#### 5 - Logs:
- a cada operação interna, é gerado um log na tabela (```currencylogs```)