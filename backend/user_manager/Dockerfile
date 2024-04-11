FROM golang:1.22.2-bullseye

WORKDIR /api

COPY ./backend/user_manager/go.mod .
COPY ./backend/user_manager/go.sum .
RUN go mod tidy

COPY ./backend/user_manager/ .
RUN go build -o main .

ENTRYPOINT [ "./main" ]