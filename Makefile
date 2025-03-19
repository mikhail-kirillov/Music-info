all: build run

tests:
	cd app; go test ./...

build:
	docker compose build

run:
	docker compose up -d

stop:
	docker compose stop

clear:
	docker compose down