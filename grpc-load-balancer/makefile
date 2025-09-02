.PHONY: all
all: run-server

.PHONY: run-server
run-server:
	@echo "Running the project server..."
	go run server/main.go

.PHONY: run-client
run-client:
	@echo "Running the project client..."
	go run client/client.go

.PHONY: proto
proto:
	@echo "Generating protobuf files..."
	protoc --go_out=. --go-grpc_out=. proto/greeter.proto
