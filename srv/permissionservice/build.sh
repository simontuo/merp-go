#!/usr/bin/env bash
cd /go/src/premissionservice && CGO_ENABLED=0 GOOS=linux GOPROXY="https://goproxy.cn,direct" go build -o ./bin/premissionservice main.go
./bin/premissionservice -dev=false