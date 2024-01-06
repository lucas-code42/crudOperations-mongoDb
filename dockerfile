# Estágio de Construção
FROM golang:1.21.4 AS BUILDER

WORKDIR /app

COPY . .

RUN go build -o crud .

EXPOSE 8080

CMD ["/app/crud"]
