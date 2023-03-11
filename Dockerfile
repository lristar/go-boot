FROM golang:alpine
WORKDIR  /app
ADD . ./
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/example/myapp cmd/example/example.go
RUN chmod +x cmd/example/myapp
ENTRYPOINT ["cmd/example/myapp"]
CMD ["-f","cmd/example/"]