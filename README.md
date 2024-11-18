# Golang Oracle Database Integration with Hexagonal Architecture

Este projeto demonstra como conectar um aplicativo Golang a um banco de dados Oracle utilizando Docker e Docker Compose. Ele segue os princípios da arquitetura hexagonal para manter o código modular e de fácil manutenção.

---

## 🚀 Recursos

- **Dockerized Environment**: Configuração do ambiente com Docker para facilitar a instalação e execução.
- **Oracle Instant Client**: Configurado para comunicação eficiente com o banco de dados Oracle.
- **Arquitetura Hexagonal**: Estrutura limpa e bem organizada para separação de responsabilidades.
- **Injeção de Dependências**: Facilita a substituição de componentes e teste unitário.

---

## 🛠️ Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- Banco de dados Oracle em execução ou acessível.

---

## 🏗️ Configuração do Ambiente

1. **Clone o repositório**
   ```bash
   git clone https://github.com/mfcbentes/go-oracle.git
   cd go-oracle
   ```
2. **Crie um arquivo .env na raiz do projeto para configurar as variáveis de ambiente necessárias, como**

   ```bash
   ORACLE_USER=seu_usuario
   ORACLE_PASSWORD=sua_senha
   ORACLE_DB=seu_db
   ORACLE_HOST=seu_host
   ORACLE_PORT=1521
   ```

3. **Inicie o Docker Compose**
   ```bash
   docker compose up --build -d
   ```
