# .env 파일에서 환경 변수 불러오기
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DOCKER_TAG := latest

## 빌드 및 실행 관련 명령어
build: ## Build production Docker images
	docker-compose build

build-dev: ## Build development Docker images without cache
	docker-compose build --no-cache

up: ## Start services with Docker Compose in detached mode
	docker-compose up -d

down: ## Stop and remove all Docker containers
	docker-compose down

restart: ## Restart Docker containers
	make down && make up

logs: ## Tail Docker Compose logs
	docker-compose logs -f

ps: ## Check status of Docker containers
	docker-compose ps
