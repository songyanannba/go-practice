#rm *.go
protoc *.proto --go_out=plugins=grpc:.
