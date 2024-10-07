# https://hub.docker.com/_/golang
FROM golang:1.23.2

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
