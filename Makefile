build:
	go build -o ./bin/go_scratch

run: build
	./bin/go_scratch

test:
	go test ./...