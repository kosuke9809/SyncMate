FROM golang:1.22.2-bullseye

WORKDIR /api

COPY ./backend/user_manage/go.mod .
COPY ./backend/user_manage/go.sum .
RUN go mod tidy

COPY ./backend/user_manage/ .
RUN go build -o main .

ENTRYPOINT [ "./main" ]