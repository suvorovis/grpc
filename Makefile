.PHONY: install-tools
install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/envoyproxy/protoc-gen-validate@v1.0.2
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1
	brew install protobuf

.PHONY: generate
generate:
	protoc --proto_path=./api/grpc \
	--go_out=. \
	--go-grpc_out=. \
	./api/grpc/**/*/*.proto