# Use a base image with Chrome installed.
FROM browserless/chrome:latest as builder

# Set the current working directory inside the container.
WORKDIR /app

# Copy the go.mod and go.sum files.
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code.
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage: Use the base image again.
FROM browserless/chrome:latest

# Copy the compiled application from the builder stage.
COPY --from=builder /app/main .

# Expose the necessary port.
EXPOSE 8080

# Command to run the executable.
CMD ["./main"]

# Start with the standard Go base image.
FROM golang:latest as builder

# Install Google Chrome.
RUN apt-get update && \
    apt-get install -y wget gnupg2 && \
    wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && \
    apt-get install -y google-chrome-stable

# Rest of your Dockerfile...
