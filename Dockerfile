# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app
COPY ./ ./

RUN go build -o ./url-shortened

EXPOSE 8080

CMD [ "/app/url-shortener", "db" ]