.PHONY: build run

build:
	go build -v ./cmd/gotify-catcher/...


run: del build
	./gotify-catcher

clean:
	go mod tidy
	go clean

start:
	./gotify-catcher

del:
	rm ./gotify-catcher || echo "file didn't exists"

.DEFAULT_GOAL := run