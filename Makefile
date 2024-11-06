run: clean build-server
	./bin/main.bin 42069

build-server:
	go build -o ./bin/main.bin ./cmd/server/main/main.go

build-script:
	go build -o ./bin/script.bin ./cmd/script/main/main.go

build: build-server build-script

server: build-server
	./bin/main.bin

script: build-script

clean:
	go clean
