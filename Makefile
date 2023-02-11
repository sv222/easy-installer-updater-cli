BINARY_NAME=easy-package-installer

all: build

build:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME) main.go

install:
	mv bin/$(BINARY_NAME) /usr/local/bin

clean:
	go clean
	rm -f /usr/local/bin/$(BINARY_NAME)