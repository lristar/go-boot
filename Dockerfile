FROM golang:alpine
WORKDIR  /app
ADD . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp cmd/example/example.go
RUN chmod +x myapp
CMD ["./myapp"]