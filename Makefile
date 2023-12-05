.PHONY: run test build

all: build run

run:

	docker-compose up

test:
	docker-compose -f ./test/docker-compose.yml down -v
	docker-compose -f ./test/docker-compose.yml up -d
	go test -count=1 ./internal/db
	docker-compose -f ./test/docker-compose.yml down -v

build:
	docker-compose build