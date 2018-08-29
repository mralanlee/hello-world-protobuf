# hello-world-protobuf # hello-world-protobuf

to compile the protobuf: `protoc -I=./protos --go_out=plugins=grpc:./protos ./protos/hello.proto`

Add a `service` in greeter and pass in a rpc with function for reply

### Go Modules
To get go modules working: `export GO111MODULE=on`

No idea how you pick versions, guess until it works LOL. Thats what I did with unknown revision errors

`go get -u` might do some upgrading dealio

go.sum is expected crypto checksums of dependencies - `go mod verify` can check these
