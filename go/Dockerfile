# https://hub.docker.com/_/golang
FROM golang:1.27.3

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
