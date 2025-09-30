#!/bin/bash

# Script para deploy no Google Cloud Run
# Uso: ./deploy.sh PROJECT_ID WEATHER_API_KEY

set -e

PROJECT_ID=$1
WEATHER_API_KEY=$2

if [ -z "$PROJECT_ID" ] || [ -z "$WEATHER_API_KEY" ]; then
    echo "Uso: $0 PROJECT_ID WEATHER_API_KEY"
    echo "Exemplo: $0 meu-projeto-123 minha-chave-api"
    exit 1
fi

echo "🚀 Iniciando deploy no Google Cloud Run..."
echo "Projeto: $PROJECT_ID"

# Configurar projeto
gcloud config set project $PROJECT_ID

# Construir e enviar imagem
echo "📦 Construindo e enviando imagem..."
gcloud builds submit --tag gcr.io/$PROJECT_ID/cloud-run-cep

# Deploy no Cloud Run
echo "🌐 Fazendo deploy no Cloud Run..."
gcloud run deploy cloud-run-cep \
  --image gcr.io/$PROJECT_ID/cloud-run-cep \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=$WEATHER_API_KEY \
  --memory 512Mi \
  --cpu 1 \
  --max-instances 10 \
  --port 8080

echo "✅ Deploy concluído!"
echo "🔗 Obtenha a URL do serviço com:"
echo "gcloud run services describe cloud-run-cep --region us-central1 --format 'value(status.url)'"
