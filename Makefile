APP?=goapp
PROJECT?=template-gin-api
VERSION?=1.0.0
PORT?=9090

run:
	go run ./cmd/main.go

clean:
	rm -rf $(APP)

test: clean
	go test -v -cover ./...

image:
	docker build . --no-cache -t $(PROJECT):$(VERSION) -f ./build/app/Dockerfile

docker:
	docker run -it --name $(PROJECT) -p $(PORT):$(PORT) --rm --env-file .env -d $(PROJECT):$(VERSION)

swagger:
	swagger generate spec -o ./docs/swagger.yaml --scan-models
	swagger serve -F=swagger ./docs/swagger.yaml

create-postgres:
	docker-compose -f ./build/postgres/docker-compose.yaml up -d

delete-postgres:
	docker-compose -f ./build/postgres/docker-compose.yaml down

create-redis:
	docker-compose -f ./build/redis/docker-compose.yaml up -d

delete-redis:
	docker-compose -f ./build/redis/docker-compose.yaml down

pre-commit:
	pre-commit install
