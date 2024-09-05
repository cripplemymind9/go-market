FROM golang:latest
WORKDIR /app
COPY . .
RUN chmod +x wait-for-it.sh
RUN go mod tidy
RUN go build -o main ./cmd/app/main.go
CMD ["./wait-for-it.sh", "postgres:5432", "--", "./main"]
