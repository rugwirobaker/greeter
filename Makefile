BINARY = bin/$(WORKINGDIR)
BUILD_FLAGS = -ldflags="-s -w" 
WORKINGDIR= $(shell pwd | rev | cut -d'/' -f 1 | rev)

build:
	env GOOS=linux CGO_ENABLED=0 go build -a -installsuffix nocgo $(BUILD_FLAGS) -o $(BINARY) ./server/.

install:
	go install ./cmd/greet

protoc:
	protoc  --go_out=plugins=grpc:. ./proto/greeter.proto

run-server:
	go run ./server/*