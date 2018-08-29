# hello-world-protobuf # hello-world-protobuf

to compile the protobuf: `protoc -I=./protos --go_out=plugins=grpc:./protos ./protos/hello.proto`

Gotta create

Add a `service` in greeter and pass in a rpc with function for reply