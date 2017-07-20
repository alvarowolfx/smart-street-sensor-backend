build:
	protoc -I api/ api/syncservice.proto --go_out=plugins=grpc:api