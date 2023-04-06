# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN go build -o main webserv.go

# Expose the port
EXPOSE 8080

# Run the Go app
CMD ["./main"]
