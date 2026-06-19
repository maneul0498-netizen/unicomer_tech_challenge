FROM golang:1.26 AS builder

WORKDIR /app

# copiar TODO el monorepo
COPY . .

# descargar deps
RUN go mod tidy

# compilar
RUN CGO_ENABLED=0 GOOS=linux go build -o unicomer_tech_challenge ./cmd/api/main.go

# ---------- RUNTIME STAGE ----------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/unicomer_tech_challenge .

CMD ["./unicomer_tech_challenge"]