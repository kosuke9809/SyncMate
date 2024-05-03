FROM golang:1.22.2-bullseye

WORKDIR /api

RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT [ "air", "-c", ".air.toml"]