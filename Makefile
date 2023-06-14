MODULE=github.com/dundee/grpc-dark-side

all: install protoc run

run:
	go run cmd/main.go

protoc: vendor
	protoc -I . -I ./vendor/github.com/cosmos/gogoproto \
		--go_opt=Mgogoproto/gogo.proto=github.com/cosmos/gogoproto/gogoproto \
		--go_opt=module=$(MODULE) \
		--go-grpc_opt=module=$(MODULE) \
		--go-grpc_out=. \
		--go_out=. \
		./proto/*.proto ./proto/nested/*.proto

protoc-node:
	cd proto/node ; yarn
	./proto/node/node_modules/grpc-tools/bin/protoc.js \
		--js_out=import_style=commonjs,binary:./proto/node/src \
		--grpc_out=./proto/node/src \
		-I . -I ./vendor/github.com/cosmos/gogoproto \
		./proto/*.proto ./proto/nested/*.proto

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

vendor:
	go mod vendor

call:
	grpcurl -plaintext 127.0.0.1:5000 foo.SomeService/GetSome

ui:
	grpcui -plaintext 127.0.0.1:5000