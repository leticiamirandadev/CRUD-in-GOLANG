# Stage 1: Build
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependências e código
COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src/
COPY docs/ ./docs/

# Compilar todo o pacote src para Linux (gerando o binário otimizado)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o crud-golang ./src

# Stage 2: Runner
FROM golang:1.26-alpine AS runner

WORKDIR /app

# Criar usuário não-root por segurança
RUN adduser -D rcode
RUN chown -R rcode:rcode /app

# Copiar o binário compilado e a documentação do Stage anterior
COPY --from=builder /app/crud-golang .
COPY --from=builder /app/docs ./docs

RUN chmod +x crud-golang

USER rcode

EXPOSE 8080

CMD ["./crud-golang"]
