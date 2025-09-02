run:
	go mod tidy
	go build -o runenv
	./runenv dev make run-check

build:
	go build -o runenv

run-check:
	./checker/check