# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app
COPY .\ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN ls -lah
RUN go build -o ./main.go

EXPOSE 8080

CMD [ "/app/URL-shortener" ]