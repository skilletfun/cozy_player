# syntax=docker/dockerfile:1
FROM python:3.11.1-slim as api

ENV UV_COMPILE_BYTECODE=1
ENV UV_PROJECT_ENVIRONMENT="/usr/local/"

WORKDIR /api

RUN apt-get update

COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/
COPY ./backend/pyproject.toml ./backend/uv.lock ./

RUN uv sync --locked

COPY ./backend .

CMD ["gunicorn", "cozy_player.wsgi", "-b", "0.0.0.0:8000", "-t", "0", "-w", "4"]


#=======================================================
FROM oven/bun:1.0 as app

WORKDIR /app

RUN apt-get update

COPY package.json bun.lock ./

RUN bun install --frozen-lockfile --production

COPY ./frontend ./

RUN bun run build

EXPOSE 3000

CMD ["bun", "run", "start"]