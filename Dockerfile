# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY ./ ./
RUN go mod download
ENV storage=""

RUN go build -o ./url-shortener

EXPOSE 8080

CMD [ "/app/url-shortener" ]