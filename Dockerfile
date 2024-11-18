# Use a imagem base do Ubuntu 22.04
FROM ubuntu:22.04

# Instalar pacotes necessários, incluindo GCC e bibliotecas de desenvolvimento
RUN apt-get update && apt-get install -y \
    libaio1 \
    gcc \
    libc6-dev \
    wget \
    unzip \
    tar \
    && rm -rf /var/lib/apt/lists/*

# Baixar e descompactar o Oracle Instant Client
RUN mkdir -p /opt/oracle && \
    wget https://download.oracle.com/otn_software/linux/instantclient/1923000/instantclient-basic-linux.x64-19.23.0.0.0dbru.zip -O /opt/oracle/instantclient.zip && \
    unzip /opt/oracle/instantclient.zip -d /opt/oracle && \
    rm /opt/oracle/instantclient.zip && \
    echo "/opt/oracle/instantclient_19_23" > /etc/ld.so.conf.d/oracle-instantclient.conf && \
    ldconfig

# Definir variáveis de ambiente para o Oracle Instant Client
ENV CGO_ENABLED=1
ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_19_23:$LD_LIBRARY_PATH
ENV PATH=/opt/oracle/instantclient_19_23:$PATH

# Instalar Go (substitua pela versão desejada)
RUN wget https://go.dev/dl/go1.23.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz && \
    rm go1.23.3.linux-amd64.tar.gz

# Definir o Go como variável de ambiente
ENV PATH=/usr/local/go/bin:$PATH

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos da aplicação para o contêiner
COPY . .

# Baixar dependências do Go
RUN go mod tidy

# Compilar o aplicativo em cmd/main.go
RUN go build -o app ./cmd/main.go

# Rodar o aplicativo
CMD ["./app"]
