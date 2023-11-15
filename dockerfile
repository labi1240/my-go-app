# Start with a base image that includes Go.
FROM golang:latest as builder

# Set the current working directory inside the container.
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker cache.
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Use a base image with Chrome installed.
FROM browserless/chrome:latest

# Copy the compiled application from the builder stage.
COPY --from=builder /app/main .

# Expose port 8080.
EXPOSE 8080

# Command to run the executable.
ENTRYPOINT ["./main"]
