.PHONY = build
build:
	mkdir -p bin
	go build -o bin/reach ./cmd/reach
