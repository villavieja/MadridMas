# MadridMas Server

The server is implemented in Go and it will run in a Google App Engine.
The buildServer.sh file should allow to compile all components.

Compilation instructions

1st step is to generate the proto generated code. 

protoc --go_out=./server/proto/ --proto_path=./server/proto/
./server/proto/madridmas.proto 

TODO(villavieja): Replace in MadridMas.pb.go
Somehow the protoc compiler still uses the old import.

- import "code.google.com/p/goprotobuf/proto"
+ import "github.com/golang/protobuf/proto"

2nd build all binaries
go build ./MadridMas/server/executable/server
go build ./MadridMas/server/executable/appengine
