# build
FROM golang:1.24.6 AS builder
WORKDIR /app

# 1) Copia solo go.mod (y go.sum si existe) para cachear deps
COPY go.mod ./
# (opcional) si ya tienes go.sum local:
# COPY go.mod go.sum ./
RUN go mod download

# 2) Copia el resto del código
COPY . .

# 3) Ajusta módulos y suma indirectas
RUN go mod tidy

# 4) Compila
RUN go build -o main .

# runtime
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
