BINARY_NAME=forwarder

all: run

run:
	go build -o $(BINARY_NAME) && $(BINARY_NAME) $(word 2, $(MAKECMDGOALS))

build:
	go build -o $(BINARY_NAME)

%:
	@: