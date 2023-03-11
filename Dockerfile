FROM golang:alpine
WORKDIR  /app
ADD . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/example/myapp cmd/example/example.go
RUN chmod +x cmd/example/myapp
CMD ["cmd/example/myapp"]