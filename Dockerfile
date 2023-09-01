FROM golang:1.20 as builder

WORKDIR /src

COPY . .

RUN go build -o /app/main ./cmd

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y curl

COPY --from=builder /app/main .

ENV PORT=8080

COPY docker.env /src/docker.env
ENV ENV_FILE_PATH=/src/docker.env

COPY  ./internal/repository/migrations /src/migrations
ENV MIGRATIONS_PATH=/src/migrations

COPY ./firebase_key.json /firebase_key.json

EXPOSE 8080

CMD ["./main"]