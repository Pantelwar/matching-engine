.PHONY: protos

protos:
	protoc -I ./ ./engine.proto --go_out=plugins=grpc:./engineGrpc