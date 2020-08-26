FROM alpine:latest

RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories

RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash

RUN apk --no-cache add ca-certificates

COPY ./build/admin /merp-go/admin

#ENTRYPOINT /gateway --registry=etcd --registry_address=172.20.0.4:2379  api --handler=httpbu

LABEL Name=merp-go-web-admin Version=0.0.1