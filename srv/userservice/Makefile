
.PHONY: proto
proto:
    
	protoc -I=./ --micro_out=./ --go_out=./  ./proto/user/user.proto

	protoc-go-inject-tag -input=./proto/user/user.pb.go


.PHONY: build
build:
	go build -o ./build/srv-user main.go

.PHONY: build-prod
build-prod:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/srv-user main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/srv-user main.go
	docker build . -t merp-go-srv-user:latest