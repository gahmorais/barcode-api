# Revisão da base e propostas de melhoria

## Melhorias implementadas neste PR

1. **Configuração via variáveis de ambiente**
   - A configuração do MongoDB e da API agora aceita variáveis de ambiente (`MONGO_URI`, `MONGO_DATABASE`, `MONGO_USER`, `MONGO_PASSWORD`, `MONGO_ADDRESS`, `MONGO_PORT`, `API_PORT`) com fallback para valores padrão.
   - Isso reduz acoplamento a ambiente local e facilita deploy em Docker/Kubernetes.

2. **Conexão com banco mais robusta**
   - `InitDb` agora usa timeout e valida conectividade com `Ping`.
   - Evita que a aplicação suba “aparentemente OK” sem conexão real com o MongoDB.

3. **Inicialização do servidor parametrizada**
   - A porta da API deixou de ser fixa e passa a ser configurável.
   - O startup log agora é sempre exibido com a porta efetiva.

## Melhorias recomendadas para próximos passos

1. **Cobertura de testes**
   - Adicionar testes unitários para controllers e repository com mocks para o Mongo.
   - Adicionar testes de integração para rotas principais (`/users`, `/products`) e fluxo de autenticação.

2. **Observabilidade**
   - Incluir middleware com request-id e logging estruturado (JSON).
   - Expor métricas Prometheus (`/metrics`) e health checks (`/healthz`, `/readyz`).

3. **Tratamento de erros e contratos de API**
   - Padronizar payload de erro HTTP em toda a API (código interno, mensagem, detalhes).
   - Diferenciar erros de validação (400), autenticação (401/403) e persistência (5xx).

4. **Segurança**
   - Garantir que JWT secret venha exclusivamente de variável de ambiente obrigatória.
   - Revisar CORS e rate limit por IP/token em endpoints sensíveis.

5. **Evolução de arquitetura**
   - Introduzir camada de service para regra de negócio, reduzindo lógica em controllers.
   - Definir interfaces para facilitar troca de implementação do repositório e testes.
