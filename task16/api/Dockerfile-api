FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o server main.go

EXPOSE 80

CMD ["/app/server"]
