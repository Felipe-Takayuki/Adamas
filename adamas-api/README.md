## ADAMAS-API

### Como rodar a API ?

```          
git clone https://github.com/Felipe-Takayuki/Adamas.git 

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

- Criação de Projeto `/repo POST`

`Authorization: Bearer JWT_TOKEN`

```json
{
  "title" : "Adamas",
  "description": "uma rede social para projetos" 
}
```

- Busca de Projeto `/search/{title} GET`


