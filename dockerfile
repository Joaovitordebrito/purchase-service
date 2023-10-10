# Use the official Golang image as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Download any necessary Go modules (if using Go modules)
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port your Go application will listen on
EXPOSE 8080

# Command to run your application
CMD ["./main"]