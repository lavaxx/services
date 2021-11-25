FROM golang:1.17-alpine3.14 As builder

ENV GO111MODULE=on \
 CGO_ENABLED=0 \
 GOOS=linux \
 GOARCH=amd64 \
 GOPROXY="https://goproxy.cn,direct"

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.7/main/ >> /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update
RUN apk add --no-cache --virtual build-dependencies \
    libffi-dev \
    openssl-dev \
    gcc \
    libc-dev \
    make

COPY . /data/src
WORKDIR /data/src
RUN make build

FROM alpine:latest

EXPOSE 8080
EXPOSE 8081

COPY --from=builder /data/src/main /data/bin/main
COPY --from=builder /data/src/.lava/ /data/.lava

WORKDIR /data

ENTRYPOINT ["/data/bin/main"]