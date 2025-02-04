run: clean build-server
	./bin/main.bin 42069

build-server:
	go build -o ./bin/main.bin ./cmd/server/main/main.go
	mkdir -p $(HOME)/.local/share/mal-cli
	cp ./bin/main.bin $(HOME)/.local/share/mal-cli/mal-cli-daemon.bin

build-script:
	go build -o ./bin/script.bin ./cmd/script/main/main.go
	cp ./bin/script.bin $(HOME)/.local/bin/mal-cli

build: build-server build-script

server: build-server
	./bin/main.bin

script: build-script
	./bin/script.bin

clean:
	go clean
