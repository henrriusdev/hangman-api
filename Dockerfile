FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy Go module files first to leverage Docker's build cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/web

# Command to run the application
CMD ["./main"]
