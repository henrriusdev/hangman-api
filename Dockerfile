FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN ls -la
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/web
CMD ["./main"]
