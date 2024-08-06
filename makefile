.PHONY: grpc


grpc:
	 protoc -I proto ./proto/gmodels/*.proto   --go_out=./proto/ --go_opt=paths=source_relative
	 protoc -I proto   --go_out=. --go-grpc_out=.  ./proto/*.proto
