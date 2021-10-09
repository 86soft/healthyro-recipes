# syntax=docker/dockerfile:1

FROM golang:1.17.1-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /healthyro-recipes

EXPOSE 80

CMD [ "/healthyro-recipes" ]