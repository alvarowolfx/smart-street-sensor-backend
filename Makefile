all: run

run:
	gom run *.go

run-grpc-test:
	go test ./grpc -bench=. -benchmem -benchtime=5s

run-http-test:
	go test ./http -bench=. -benchmem -benchtime=5s

build:
	GOOS=linux GOARCH=amd64 gom build -ldflags '-s'

fmt: 
	gom fmt .