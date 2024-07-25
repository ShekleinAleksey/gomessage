# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /gomessage

COPY . /gomessage

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o main main.go

EXPOSE 8080

# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gomessage

CMD ["main"]