BINARY_NAME=umm
BINARY_PATH=./bin/$(BINARY_NAME)

build:
	@go build -o $(BINARY_PATH) main.go 

install: 
	@go install -v ./...

clean:
	@rm -rf ./bin 
