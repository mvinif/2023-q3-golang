build:
	docker build . -t golang-api:latest

run-dev:
	docker compose -f docker-compose-dev.yml up -d

run-dev2:
	docker compose -f docker-compose-dev.yml up

stop-dev:
	docker compose -f docker-compose-dev.yml down

reload-dev:
	docker compose -f docker-compose-dev.yml down
	make build
	docker compose -f docker-compose-dev.yml up -d

watch-dev:
	docker compose -f docker-compose-dev.yml watch

up:
	docker compose up -d

down:
	docker compose down

test:
	go test ./...
