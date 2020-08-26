.PHONY: build
build:
	go build -o ./build/admin main.go

.PHONY: build-prod
build-prod:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/admin main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/admin main.go
	docker build . -t merp-go-web-admin:latest