FROM golang:1.23-alpine

WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy everything else
COPY . .

# Debug: Check the contents of the /app directory
RUN ls -la /app && ls -la /app/cmd/web

# Attempt to build
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/web

CMD ["./main"]
