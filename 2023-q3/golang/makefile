build:
	docker build . -t golang-api:latest

run-dev:
	docker compose -f docker-compose-dev.yml up -d --build

run-dev2:
	docker compose -f docker-compose-dev.yml up --build

stop-dev:
	docker compose -f docker-compose-dev.yml down

reload-dev:
	docker compose -f docker-compose-dev.yml down
	make build
	docker compose -f docker-compose-dev.yml up -d --build

watch-dev:
	docker compose -f docker-compose-dev.yml watch

up:
	docker compose up -d

down:
	docker compose down

test:
	go test ./...

# $ make go CMD="build test"
go:
	docker run --rm -v "$$PWD:/app" -w /app golang go $(CMD)
