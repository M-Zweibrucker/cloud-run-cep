# Cloud Run CEP

URL NA CLOUD RUN : https://avaliacao-pos-709409355709.us-central1.run.app/weather/01310-100

Sistema em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). Publicado no Google Cloud Run.

## Funcionalidades

- ✅ Validação de CEP (8 dígitos)
- ✅ Consulta de localização via ViaCEP API
- ✅ Consulta de clima via WeatherAPI
- ✅ Conversão de temperaturas (Celsius, Fahrenheit, Kelvin)
- ✅ Tratamento de erros adequado
- ✅ Testes automatizados
- ✅ Containerização com Docker
- ✅ Deploy no Google Cloud Run

## Requisitos

- Go 1.21+
- Docker e Docker Compose
- Chave da API WeatherAPI (gratuita)

## Configuração

1. Clone o repositório
2. Copie o arquivo de exemplo de ambiente:
   ```bash
   cp env.example .env
   ```

3. Configure sua chave da WeatherAPI no arquivo `.env`:
   ```
   WEATHER_API_KEY=sua_chave_aqui
   ```

## Executando Localmente

### Com Docker Compose (Recomendado)

```bash
# Construir e executar
docker-compose up --build

# Em modo detached
docker-compose up -d --build
```

### Com Go diretamente

```bash
# Instalar dependências
go mod download

# Executar testes
go test ./...

# Executar aplicação
go run main.go
```

## Testando a API

### Sucesso (CEP válido)
```bash
curl http://localhost:8080/weather/01310-100
```

Resposta esperada:
```json
{
  "temp_C": 25.0,
  "temp_F": 77.0,
  "temp_K": 298.0
}
```

### CEP inválido (formato incorreto)
```bash
curl http://localhost:8080/weather/1234567
```

Resposta esperada:
```json
{
  "error": "invalid zipcode"
}
```

### CEP não encontrado
```bash
curl http://localhost:8080/weather/99999999
```

Resposta esperada:
```json
{
  "error": "can not find zipcode"
}
```

## Deploy no Google Cloud Run

### Pré-requisitos

1. Instalar Google Cloud CLI
2. Configurar projeto e autenticação:
   ```bash
   gcloud auth login
   gcloud config set project SEU_PROJECT_ID
   ```

### Deploy

```bash
# Construir e enviar para Container Registry
gcloud builds submit --tag gcr.io/SEU_PROJECT_ID/cloud-run-cep

# Deploy no Cloud Run
gcloud run deploy cloud-run-cep \
  --image gcr.io/SEU_PROJECT_ID/cloud-run-cep \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=sua_chave_aqui
```

### Testando o Deploy

Após o deploy, você receberá uma URL. Teste com:

```bash
curl https://sua-url-do-cloud-run.run.app/weather/01310-100
```

## Estrutura do Projeto

```
cloud-run-cep/
├── main.go                    # Ponto de entrada da aplicação
├── go.mod                     # Dependências Go
├── Dockerfile                 # Configuração Docker
├── docker-compose.yml         # Configuração Docker Compose
├── env.example               # Exemplo de variáveis de ambiente
├── README.md                 # Este arquivo
└── internal/
    ├── handlers/             # Handlers HTTP
    │   ├── weather_handler.go
    │   └── handlers_test.go
    └── services/             # Serviços de negócio
        ├── cep_service.go
        ├── weather_service.go
        └── services_test.go
```

## APIs Utilizadas

- **ViaCEP**: https://viacep.com.br/ - Consulta de CEP
- **WeatherAPI**: https://www.weatherapi.com/ - Consulta de clima

## Conversões de Temperatura

- **Fahrenheit**: F = C × 1.8 + 32
- **Kelvin**: K = C + 273

## Códigos de Resposta HTTP

- **200**: Sucesso - Temperaturas retornadas
- **404**: CEP não encontrado
- **422**: CEP com formato inválido
- **500**: Erro interno do servidor
