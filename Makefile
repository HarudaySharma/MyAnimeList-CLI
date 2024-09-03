run: clean build-server
	./bin/main.bin

build-server:
	go build -o ./bin/main.bin ./cmd/server/main.go

build-script:
	go build -o ./bin/script.bin ./cmd/script/main.go

build: build-server build-script

script: build-script
	./bin/script.bin

clean:
	go clean 
