
# API Busca CEP

Este é um projeto simples em Go para buscar informações de endereços a partir de um CEP usando a API do BrasilAPI.


## Pré-requisitos
Certifique-se de ter o Go instalado em sua máquina.

## Rodar o projeto

Instalação das dependências

```bash
  go mod tidy
```


Executar o projeto

```bash
  go run .
```



## Documentação da API

#### Request GET

```http
  GET /{cep}
```

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `cep` | `number` | CEP em formato numérico |



Exemplo de resposta:

```
{
  "cep": "01001000",
  "state": "SP",
  "city": "São Paulo",
  "neighborhood": "Sé",
  "street": "Praça da Sé",
  "service": "brasilapi",
  "location": {
    "type": "Point",
    "coordinates": {
      "longitude": "-46.6333",
      "latitude": "-23.5503"
    }
  }
}
