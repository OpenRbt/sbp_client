FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/main ./cmd

FROM alpine

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

RUN apk update --no-cache && apk add --no-cache ca-certificates

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /app/main .

# Копируем миграции
COPY ./internal/repository/migrations /src/migrations
ENV MIGRATIONS_PATH=/src/migrations

# Копируем файл ключа Firebase
COPY ./firebase_key.json .
ENV FB_KEYFILE_PATH=/firebase_key.json

# Копируем файл с переменными окружения
COPY docker.env /src/docker.env
ENV ENV_FILE_PATH=/src/docker.env

EXPOSE 8080

CMD ["./main"]