all: run

tidy:
	@if [ ! -f go.sum ]; then \
		go mod tidy; \
	fi

run: tidy
	go run ./backend
