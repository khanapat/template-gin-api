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

build-container:
	docker build . --no-cache -t $(PROJECT):$(VERSION) -f ./build/Dockerfile

run-container:
	docker run --env-file .env -d $(PROJECT):$(VERSION)