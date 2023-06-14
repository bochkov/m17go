.PHONY: all deps release run clean

all: run

deps:
	go mod tidy

release:
	go build -o m17 -ldflags "-s -w"

run:
	go run .
	
clean:
	go clean -cache
	rm -f m17
