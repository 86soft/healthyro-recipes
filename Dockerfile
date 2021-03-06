# syntax=docker/dockerfile:1
FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" .

FROM scratch

WORKDIR /app

COPY --from=builder /app/healthyro-recipes /usr/bin/

EXPOSE 80

ENTRYPOINT ["healthyro-recipes"]