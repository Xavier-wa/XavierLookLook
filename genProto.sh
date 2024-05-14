protoc -I $(pwd) --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. $1.proto
# echo $(pwd)