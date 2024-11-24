build:
	docker build . -t golang-api:latest

run_dev:
	docker compose -f docker-compose-dev.yml up -d

stop_dev:
	docker compose -f docker-compose-dev.yml down

reload_dev:
	docker compose -f docker-compose-dev.yml down
	make build
	docker compose -f docker-compose-dev.yml up -d

watch_dev:
	docker compose -f docker-compose-dev.yml watch

up:
	docker compose up -d

down:
	docker compose down

test:
	go test ./...
