# syntax=docker/dockerfile:1
FROM python:3.11.1-slim as api

ENV UV_COMPILE_BYTECODE=1
ENV UV_PROJECT_ENVIRONMENT="/usr/local/"
ENV MUSIC_FOLDER="/music" 

WORKDIR /api

RUN apt-get update

COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/
COPY ./backend/pyproject.toml ./backend/uv.lock ./

RUN uv sync --locked

COPY ./backend .


#=======================================================
FROM oven/bun:1.2.8 as build-app

ARG API_URL

WORKDIR /app

RUN apt-get update

COPY ./frontend/package.json ./frontend/bun.lock ./

RUN bun install --frozen-lockfile

COPY ./frontend ./
RUN echo 'PUBLIC_API_URL="${API_URL}"' > .env

RUN bun run build


#=======================================================
FROM nginx:alpine AS app

COPY .nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build-app /app/build /usr/share/nginx/html

ENTRYPOINT ["nginx", "-g", "daemon off;"]
