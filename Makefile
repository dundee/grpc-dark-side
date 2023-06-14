all: install protoc run

run:
	go run cmd/main.go

protoc:
	protoc -I . \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_out=. \
		--go_out=. \
		./proto/*.proto ./proto/nested/*.proto

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

call:
	grpcurl -plaintext -proto ./proto/some.proto 127.0.0.1:5000 foo.SomeService/GetSome

ui:
	grpcui -plaintext 127.0.0.1:5000