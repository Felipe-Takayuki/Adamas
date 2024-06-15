## ADAMAS-API

### Como rodar a API ?

```          
git clone https://github.com/Adamas-Projects/Adamas-BackEnd.git 

cd Adamas/adamas-api

docker-compose up --build
```
OU 
```
git clone https://github.com/Felipe-Takayuki/Adamas.git 

cd Adamas/adamas-api/cmd/api

go run .
``` 
### E iniciar o banco de dados manualmente
##

```http://localhost:3000```

- Cadastro de Usuários `/create POST`

corpo da requisição:  

```json
{ 
  "name" : "felipe-takayuki",
  "email" : "felipe@gmail.com",
  "password" : "felipe123" // a senha é criptografada
}
/// a requisição retorna um token jwt que será usado em outros endpoints 
```

- Login de Usuários `/login POST`
```json
{ 
  "email" : "felipe@gmail.com",
  "password" : "felipe123" // a senha é criptografada
}
/// a requisição retorna um token jwt que será usado em outros endpoints 
```
- Cadastro de Instituição `/create_institution POST`
```json
{
    "name" : "ETEC",
    "email": "etec@etec.com",
    "password": "12345678", // a senha é criptografada
    "cnpj": 28301041000137
}
/// a requisição retorna um token jwt que será usado em outros endpoints
```
- Login de Instituição `/login_institution POST`
```json
{
    "email": "etec@etec.com",
    "password": "12345678" // a senha é criptografada
}
/// a requisição retorna um token jwt que será usado em outros endpoints
```

- Criação de Projeto `/repo POST`

`Authorization: Bearer JWT_TOKEN`

```json
{
  "title" : "Adamas",
  "description": "uma rede social para projetos" ,
  "content": "### Olá" // deve ser em markdown
}
```

- Busca de Projeto `/repo/{title} GET`

- Obter Projetos `/repo GET`

- Criação do Evento `/event POST`

`Authorization: Bearer JWT_TOKEN`
```json 
{
    "name": "Amostra de TCC",
    "address" : "ETEC ANTONIO DEVISATE, avenida castro alves",
    "date": "2020-12-02",
    "description": "Uma amostra de tcc uai"
}
