.PHONY: proto
proto:
    
	protoc -I=./ --micro_out=./ --go_out=./  ./proto/sms/sms.proto

	protoc-go-inject-tag -input=./proto/sms/sms.pb.go

.PHONY: build
build:
	go build -o ./build/srv-sms main.go

.PHONY: build-prod
build-prod:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/srv-sms main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/srv-sms main.go
	docker build . -t merp-go-srv-sms:latest

