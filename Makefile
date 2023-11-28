.PHONY: all deps release run clean

all: run

deps:
	go mod tidy

release:
	go build -o m17 -ldflags "-s -w" ./cmd/main.go

run:
	go run ./cmd/main.go
	
clean:
	go clean -cache
	rm -f m17
