all: rebuild
test: lint run-tests

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: build
build:
	docker build --target bin --output bin/ .

.PHONY: rebuild
rebuild: clean build

.PHONY: run-tests
run-tests:
	docker build --target test .

.PHONY: lint
lint:
	docker build --target lint .

.PHONY: db
db:
	docker run --name=film-db -e POSTGRES_PASSWORD="film-db" -d --rm postgres:16.2

.PHONY: launch
launch:
	docker build -t film_db .
