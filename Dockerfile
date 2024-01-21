FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache protoc protobuf-dev

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN go mod download

RUN protoc -Iproto --go_opt=module=vocab_service --go_out=. --go-grpc_opt=module=vocab_service --go-grpc_out=. proto/*.proto

RUN go build -o bin/vocab_service ./cmd/.

EXPOSE 50052

CMD ["./bin/vocab_service"]