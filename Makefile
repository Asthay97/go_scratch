build:
	go build -o ./bin/blockchain_go

run: build
	./bin/blockchain_go

test:
	go test ./...