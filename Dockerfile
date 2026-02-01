# syntax=docker/dockerfile:1
FROM golang:1.24-alpine as build-api

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /api

COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download

COPY ./backend .

RUN go build -o main ./cmd/main/main.go


#=======================================================
FROM oven/bun:1.2.8 as build-app

WORKDIR /app

COPY ./frontend/package.json ./frontend/bun.lock ./

RUN bun install --frozen-lockfile

COPY ./frontend ./
RUN echo 'PUBLIC_API_URL="/api"' > .env

RUN bun run build


#=======================================================
FROM nginx:alpine AS app

COPY --from=build-api /api/main /app/main

COPY .nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build-app /app/build /usr/share/nginx/html

COPY ./run.sh run.sh
RUN chmod +x run.sh

CMD ./run.sh

EXPOSE 80
