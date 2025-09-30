#!/bin/bash

# Script de teste para demonstrar o funcionamento da API
# Uso: ./test.sh [URL_BASE]

BASE_URL=${1:-"http://localhost:8080"}

echo "🧪 Testando API Cloud Run CEP"
echo "URL Base: $BASE_URL"
echo ""

# Teste 1: CEP válido
echo "Teste 1: CEP válido (01310-100 - São Paulo)"
curl -s "$BASE_URL/weather/01310-100" | jq '.' 2>/dev/null || curl -s "$BASE_URL/weather/01310-100"
echo ""
echo ""

# Teste 2: CEP inválido (formato incorreto)
echo "Teste 2: CEP inválido (formato incorreto)"
curl -s "$BASE_URL/weather/1234567" | jq '.' 2>/dev/null || curl -s "$BASE_URL/weather/1234567"
echo ""
echo ""

# Teste 3: CEP não encontrado
echo "Teste 3: CEP não encontrado"
curl -s "$BASE_URL/weather/99999999" | jq '.' 2>/dev/null || curl -s "$BASE_URL/weather/99999999"
echo ""
echo ""

# Teste 4: CEP com formatação diferente
echo "Teste 4: CEP com formatação diferente (22071-900 - Rio de Janeiro)"
curl -s "$BASE_URL/weather/22071900" | jq '.' 2>/dev/null || curl -s "$BASE_URL/weather/22071900"
echo ""
echo ""

echo "✅ Testes concluídos!"
