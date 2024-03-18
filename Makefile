all: build launch
test: lint run-tests

.PHONY: build
build:
	docker build -t film_db .

.PHONY: run-tests
run-tests:
	docker build --target test .

.PHONY: lint
lint:
	docker build --target lint .

.PHONY: launch
launch:
	docker compose up
