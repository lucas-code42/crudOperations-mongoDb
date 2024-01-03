# Estágio de Construção
FROM golang:1.19 AS BUILDER

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o crud .

# Estágio de Execução
FROM alpine:3.15

WORKDIR /app

RUN adduser -D testUser

COPY --from=BUILDER /app/crud /app/crud

RUN chown -R testUser:testUser /app
RUN chmod +x /app/crud

USER testUser

EXPOSE 8080

CMD ["/app/crud"]
