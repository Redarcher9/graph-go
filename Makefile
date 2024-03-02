start:
	go run cmd/http/main.go

graphgen:
	go run github.com/99designs/gqlgen generate

protogen:
	protoc --go_out=. --go-grpc_out=. proto/greet.proto