build:
	docker build . -t golang-api:latest

run_dev:
	docker compose -f docker-compose-dev.yaml up -d

stop_dev:
	docker compose -f docker-compose-dev.yaml down

reload_dev:
	docker compose -f docker-compose-dev.yaml down
	make build
	docker compose -f docker-compose-dev.yaml up -d

up:
	docker compose up -d

down:
	docker compose down

test:
	go test ./...
