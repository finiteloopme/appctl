dummy:

protoc:
	protoc --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		proto/app.proto

tidy:
	go mod tidy

build: tidy
	go build .

test:

run:
	go run main.go

release: