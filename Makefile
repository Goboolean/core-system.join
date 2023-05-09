grpc-generate:
	protoc -I api/grpc/ \
		--go_out=api/grpc/	\
		--go_opt=paths=source_relative \
		--go-grpc_out=api/grpc/	\
		--go-grpc_opt=paths=source_relative	\
		api/grpc/query-server.proto
