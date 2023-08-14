generate-proto:
	protoc --go_out=$2 --go_opt=paths=source_relative \
		--go-grpc_out=$2 --go-grpc_opt=paths=source_relative \
		$1
