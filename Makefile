all: test lint vet build

build: server client

server:
	@cd cmd/server && go build -o ../../bin/server

client:
	@cd cmd/client && go build -o ../../bin/client

lint:
	@echo "Linting..."
	@golint ./...

vet:
	@echo "Vetting..."
	@go vet ./...

test:
	@echo "Testing..."
	@go test ./...

clean:
	@find . -name "*-wal" -delete
	@find . -name "*-shm" -delete
	@rm -f bin/*.linux

gen:
	@protolint api/proto/
	@go generate ./...

deps: dep_protoc dep_protolint

dep_protoc:
	@go get -u github.com/golang/protobuf/protoc-gen-go

dep_protolint:
	@go get -u -v github.com/yoheimuta/protolint/cmd/protolint
