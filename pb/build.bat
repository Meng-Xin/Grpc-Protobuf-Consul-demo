protoc --go_out=. --go-grpc_out=.  *.proto
@REM protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./person.proto

