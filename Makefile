Domain=services
Base=github.com/pubgo/lava
Version=$(shell git tag --sort=committerdate | tail -n 1)
BuildTime=$(shell date "+%F %T")
CommitID=$(shell git rev-parse --short=8 HEAD)
LDFLAGS=-ldflags " \
-X '${Base}/version.BuildTime=${BuildTime}' \
-X '${Base}/version.CommitID=${CommitID}' \
-X '${Base}/version.Version=${Version}' \
-X '${Base}/version.Domain=${Domain}' \
-X '${Base}/version.Data=hello' \
"

.PHONY: build
build:
	go build ${LDFLAGS} -v -o main cmd/*.go

.PHONY: test./ma
test:
	@go test -short -race -v ./... -cover

ci:
	@golangci-lint run -v --timeout=5m

proto-gen:
	debug=1 lava protoc gen

proto-rm:
	@rm -rf proto/**/*.go

proto-vendor:
	lava protoc vendor
	lava protoc ls

cover:
	gocov test -tags "kcp quic" ./... | gocov-html > cover.html
	open cover.html

vet:
	@go vet ./...

lint:
	@golangci-lint run --skip-dirs-use-default --timeout 3m0s
