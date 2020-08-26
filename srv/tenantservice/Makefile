
.PHONY: proto
proto:
    
	protoc -I=./ --micro_out=./ --go_out=./  ./proto/tenant/tenant.proto

	protoc-go-inject-tag -input=./proto/tenant/tenant.pb.go

	protoc -I=./ --micro_out=./ --go_out=./  ./proto/department/department.proto

	protoc-go-inject-tag -input=./proto/department/department.pb.go

	protoc -I=./ --micro_out=./ --go_out=./  ./proto/tenantuser/tenantuser.proto

	protoc-go-inject-tag -input=./proto/tenantuser/tenantuser.pb.go


.PHONY: build
build:
	go build -o ./build/srv-tenant main.go


.PHONY: build-prod
build-prod:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/srv-tenant main.go


.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/srv-tenant main.go
	docker build . -t merp-go-srv-tenant:latest

    