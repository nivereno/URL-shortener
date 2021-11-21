# syntax=docker/dockerfile:1

FROM golang:latest

RUN mkdir /app
COPY ./ /app

RUN go mod download

RUN

Expose 8080