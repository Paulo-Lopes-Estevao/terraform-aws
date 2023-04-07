gen:
	protoc --proto_path=app/grpc/protofile app/grpc/protofile/messageuser.proto --go_out=app/grpc/pb --go-grpc_out=app/grpc/pb

test:
	go test -v ./...

test coverage:
	go test -v ./... -coverprofile=coverage.out