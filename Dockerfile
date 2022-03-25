FROM golang:1.18-alpine

RUN apk add git

COPY . /app

WORKDIR /app/cmd

RUN go build -o server

EXPOSE 8080

CMD ./server
# CMD ./server -conf illia.conf