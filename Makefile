
build:
	protoc -Iproto --go_opt=module=vocab_service --go_out=. --go-grpc_opt=module=vocab_service --go-grpc_out=. proto/*.proto
	go build -o bin/vocab_service.exe ./cmd/.