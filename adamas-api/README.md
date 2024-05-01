## ADAMAS-API

### Como rodar a API ?

```          
git clone https://github.com/Felipe-Takayuki/Adamas.git 

cd Adamas/adamas-api

docker-compose up --build
```


```http://localhost:8080```

- Cadastro de Usuários `/create POST`

corpo da requisição:  

```json
{ 
  "name" : "felipe-takayuki",
  "email" : "felipe@gmail.com",
  "password" : "felipe123" // a senha é criptografada
}
```

- Login de Usuários `/login POST`
```json
{ 
  "email" : "felipe@gmail.com",
  "password" : "felipe123" // a senha é criptografada
}
```