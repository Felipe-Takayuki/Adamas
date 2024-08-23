# Documentação da API

## Endpoints

### Criar Usuário

**POST** `http://localhost:3000/create`

```json
{
    "name": "Felipe",
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

**POST** `http://localhost:3000/project`

**Authorization**: Bearer TOKEN

```json
{
    "title": "Adamas",
    "description": "uma rede social para projetos",
    "content": "###hello world!"
}

```

### Atualizar Projeto

**PUT** `http://localhost:3000/project/{project_id}`

**Authorization**: Bearer TOKEN

```json
{
    "title": "Adamas-Projects",
    "description": "uma rede social para a divulgação de projetos e eventos",
    "content": "###hello world 2"
}

```

### Deletar Projeto

**DELETE** `http://localhost:3000/project/{project_id}`

**Authorization**: Bearer TOKEN

```json
{
    "email": "felipe@email.com",
    "password": "12345678"
}

```

### Adicionar Categoria ao Projeto

**POST** `http://localhost:3000/project/{project_id}/category`

**Authorization**: Bearer TOKEN

```json
{
    "category_name": "ti"
}

```

### Adicionar Comentário ao Projeto

**POST** `http://localhost:3000/project/{project_id}/comment`

**Authorization**: Bearer TOKEN

```json
{
    "comment": "muito brabo"
}

```

### Deletar Comentário do Projeto

**DELETE** `http://localhost:3000/project/{project_id}/comment`

**Authorization**: Bearer TOKEN

```json
{
    "comment_id": 1
}

```

### Atualizar Comentário do Projeto

**PUT** `http://localhost:3000/project/{project_id}/comment`

**Authorization**: Bearer TOKEN

```json
{
    "comment_id": 1,
    "comment": "muito brabo, eu sei"
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

### Adicionar Usuário ao Projeto

**POST** `http://localhost:3000/project/{project_id}/add-user`

**Authorization**: Bearer TOKEN

```json
{
    "user_id": 2
}

```

### Atualizar Evento

**PUT** `http://localhost:3000/event/{event_id}`

**Authorization**: Bearer TOKEN

```json
{
    "name": "Amostra de TCC 001",
    "description": "Uma amostra de tcc, com projetos dos alunos da ETEC"
}

```

### Adicionar Sala ao Evento

**POST** `http://localhost:3000/event/{event_id}/room`

**Authorization**: Bearer TOKEN

```json
{
    "name": "Sala 5",
    "quantity_projects": 10
}

```

### Inscrever-se no Evento

**POST** `http://localhost:3000/event/{event_id}/subscribe`

**Authorization**: Bearer TOKEN

### Listar Inscritos no Evento

**GET** `http://localhost:3000/event/{event_id}/subscribers`

**Authorization**: Bearer TOKEN

### Participação no Evento

**POST** `http://localhost:3000/event/{event_id}/participation`

**Authorization**: Bearer TOKEN

```json
{
    "project_id": 1
}

```

### Aprovar Participação no Evento

**POST** `http://localhost:3000/event/{event_id}/approve-participation`

**Authorization**: Bearer TOKEN

```json
{
    "project_id": 1,
    "room_id": 1
}

```

### Pesquisar Projetos

**GET** `http://localhost:3000/project/search`

### Pesquisar Projeto Específico

**GET** `http://localhost:3000/project/search/{project_title}`

### Listar Projetos do Usuário

**GET** `http://localhost:3000/project/user/{user_id}`

### Pesquisar Evento Específico

**GET** `http://localhost:3000/event/search/{event_title}`