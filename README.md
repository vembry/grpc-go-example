## grpc examples on go-server go-client

### structure
```
./
├─ server/
│   ├─ main.go
│   └─ chat.proto
├─ proto/
│   └─ auto generated protobuf files
└─ client/
    └─ main.go
```
- `server` is serving grpc
- `client` is calling `server` rpc
- `proto` is providing protobuf files for `client` to use

### generate protobuf
> protoc --go_out=./proto --go-grpc_out=./proto server/chat.proto

### how to run
services have to run separately, assuming you are on the project's root

server
> go run server/main.go

client
> go run client/main.go