export GOFLAGS=-mod=vendor
APPCTL_CLI=appctl
OUTPUT_DIR=bin
dummy:

code-gen:
	go generate ./pkg/... ./cmd/...

protoc:
	protoc --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		proto/app.proto

ensure-deps:
	go mod tidy
	go mod vendor

build: ensure-deps
	go build -o ${OUTPUT_DIR}/${APPCTL_CLI} ./cmd/...

test:

run: ensure-deps
	go run cmd/appctl.go

release: