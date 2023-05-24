FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main webserv.go
EXPOSE 8080
CMD ["./main"]

