# Application build & run
build: 
	go build -o main.o command/main.go
run: 
	go run command/main.go

# Proto compile for message & service
rpc-msg:
	protoc --go_out=internal/adapters/framework/left/grpc \
	--proto_path=internal/adapters/framework/left/grpc/proto \
	internal/adapters/framework/left/grpc/proto/*_msg.proto
rpc-svc:
	protoc --go-grpc_out=internal/adapters/framework/left/grpc \
	--proto_path=internal/adapters/framework/left/grpc/proto \
	internal/adapters/framework/left/grpc/proto/*_svc.proto
rpc-svc-unsafe:
	protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/left/grpc \
	--proto_path=internal/adapters/framework/left/grpc/proto \
	internal/adapters/framework/left/grpc/proto/*_svc.proto
rpc-buil-rm:
	rm -rf internal/adapters/framework/left/grpc/proto/