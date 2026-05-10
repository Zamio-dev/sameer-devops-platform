# Sameer DevOps Platform — Root Makefile
# This is your command center. Run `make help` to see all commands.

.PHONY: help up down build logs ps clean setup-dev

# Default target — runs when you type just `make`
help:
	@echo ""
	@echo "Sameer DevOps Platform"
	@echo "======================"
	@echo "make up        — Start all services"
	@echo "make down      — Stop all services"
	@echo "make build     — Rebuild all containers"
	@echo "make logs      — Tail all logs"
	@echo "make ps        — Show running containers"
	@echo "make clean     — Remove containers, volumes, networks"
	@echo "make setup-dev — Install all dev tools"
	@echo ""

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose build --no-cache

logs:
	docker compose logs -f

ps:
	docker compose ps

clean:
	docker compose down -v --remove-orphans

setup-dev:
	@echo "Installing development tools..."
	@./scripts/setup-dev.sh
