include .env
SHELL := bash

up:
	@(docker-compose up -d)

ps:
	@(docker-compose ps)

migrate:
	goose -dir ./migrations/postgres postgres "postgres://${DATABASE_NAME}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?search_path=${DATABASE_SCHEMA}" up

down:
	@(docker-compose down -v --remove-orphans)
