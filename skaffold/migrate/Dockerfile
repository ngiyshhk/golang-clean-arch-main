FROM golang:1.12-alpine
RUN apk add git gcc libc-dev
COPY dbconfig.yml dbconfig.yml
COPY migrations migrations
RUN go get -v github.com/rubenv/sql-migrate/...
