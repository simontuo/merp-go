.PHONY: proto
proto:
    
	protoc -I=./ --micro_out=./ --go_out=./  ./proto/role/role.proto

	protoc-go-inject-tag -input=./proto/role/role.pb.go

	protoc -I=./ --micro_out=./ --go_out=./  ./proto/role/roleuser.proto
	
	protoc-go-inject-tag -input=./proto/role/roleuser.pb.go

	protoc -I=./ --micro_out=./ --go_out=./  ./proto/permission/permission.proto

	protoc-go-inject-tag -input=./proto/permission/permission.pb.go

	protoc -I=./ --micro_out=./ --go_out=./  ./proto/permission/permissionrole.proto

	protoc-go-inject-tag -input=./proto/permission/permissionrole.pb.go


.PHONY: build
build:
	go build -o ./build/srv-permission main.go

.PHONY: build-prod
build-prod:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/srv-permission main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/srv-permission main.go
	docker build . -t merp-go-srv-permission:latest

