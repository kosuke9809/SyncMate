FROM golang:1.22.2-bullseye

WORKDIR /api

COPY ./backend/user_management_service/go.mod .
COPY ./backend/user_management_service/go.sum .
RUN go mod tidy

COPY ./backend/user_management_service/ .
RUN go build -o main .

ENTRYPOINT [ "./main" ]