
# Como rodar a API ?

```          
git clone https://github.com/Adamas-Projects/Adamas-BackEnd.git

cd Adamas/adamas-api

docker-compose up --build
```
OU 
```
git clone https://github.com/Adamas-Projects/Adamas-BackEnd.git

cd Adamas-BackEnd/adamas-api/cmd/api

go run .
``` 
### E iniciar o banco de dados manualmente
##


# Documentação da API

## Endpoints

### Criar Usuário

**POST** `http://localhost:3000/create`

```json
{
    "name": "Felipe321",
    "nickname": "felipe",
    "description": "o brabo", //não é obrigatório
    "email": "felipe@email.com",
    "password": "12345678"
}

```

### Login de Usuário

**POST** `http://localhost:3000/login`

```json
{
    "email": "felipe@email.com",
    "password": "12345678"
}

```

### Criar Instituição

**POST** `http://localhost:3000/create/institution`

```json
{
    "name": "ETEC",
    "email": "etec@etec.com",
    "password": "12345678",
    "cnpj": "28301041000137"
}

```

### Login de Instituição

**POST** `http://localhost:3000/login/institution`

```json
{
    "email": "etec@etec.com",
    "password": "12345678"
}

```

### Criar Projeto

**POST** `http://localhost:3000/project`**Authorization**: Bearer TOKEN

```json
{
    "title": "Adamas",
    "description": "uma rede social para projetos",
    "content": "###hello world!"
}

```

### Atualizar Projeto

**PUT** `http://localhost:3000/project/1`

**Authorization**: Bearer TOKEN

```json
{
    "title": "Adamas-Projects",
    "description": "uma rede social para a divulgação de projetos e eventos",
    "content": "###hello world 2"
}

```

### Deletar Projeto

**DELETE** `http://localhost:3000/project/1`

**Authorization**: Bearer TOKEN

```json
{
    "email": "felipe@email.com",
    "password": "12345678"
}

```

### Adicionar Categoria ao Projeto

**POST** `http://localhost:3000/project/1/category`

**Authorization**: Bearer TOKEN

```json
{
    "repository_id": 1,
    "category_name": "ti"
}

```

### Adicionar Comentário ao Projeto

**POST** `http://localhost:3000/project/1/comment`

**Authorization**: Bearer TOKEN

```json
{
    "comment": "muito brabo"
}

```

### Deletar Comentário do Projeto

**DELETE** `http://localhost:3000/project/1/comment`

**Authorization**: Bearer TOKEN

```json
{
    "comment_id": 1
}

```

### Criar Evento

**POST** `http://localhost:3000/event`

**Authorization**: Bearer TOKEN

```json
{
    "name": "Amostra de TCC",
    "address": "ETEC ANTONIO DEVISATE, avenida castro alves",
    "date": "2020-12-02",
    "description": "Uma amostra de tcc uai"
}

```

### Inscrição no Evento

**POST** `http://localhost:3000/event/subscribe/1`

**Authorization**: Bearer TOKEN

### Listar Projetos

**GET** `http://localhost:3000/project`

### Obter Projeto Específico

**GET** `http://localhost:3000/project/Adamas`

### Obter Evento Específico

**GET** `http://localhost:3000/event/Amostra de TCC`