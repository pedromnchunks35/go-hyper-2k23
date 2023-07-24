# Introduction to gRPC
- It is a protocol that uses protobuf, to call a object from a GRPC Server, that by using protobuf enables the client to run it like it is a local object
- As another RPCS, it is based on defining a certain service using interfaces, mentioning the methods
- The server implements its interface and runs a GRPC server to handle the client calls
- On the client side, the client has a stub, that provides the same methods as the server
![Grpc representation](assets/grpc.png)
- Servers and clients can be writen in different programming languages
- gRPC uses [Protocol Buffers](Protocol-buffers/readme.md) by default, which is a different type of serialization of data such as JSON (we can still use json)

##  Quick Start
### Dependencies
- go
- protoc (protocol buffer compiler)
- go plugins
  ```
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
  ```
### Run Hello world code [Link](quickstart/grpc-go/examples/helloworld/)
We made a edit to the proto file inside of examples/helloworld/helloworld and we run:
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```
Run both programs once again

# Core concepts, architecture and lifecycle
- It lets you define four types of services
  ```
    (Unary RPCS) Single request and receive single response back   
        -> rpc SayHello(HelloRequest) returns (HelloResponse);
    
    (Server streaming RPCs) Client sends a request to the server and gets a stream to read a sequence of messages back. Client reads from the returned stream until there are no more messages. gRPC guarantees message ordering within an individual RPC call. gRPC garantees message ordering within an individual RPC call.
        -> rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
    
    (Client streaming RPCs) Client writes a stream of messages and sends them to the server. Client waits until the server reads all the messages

    (Bidirectional Streaming RPCs) Streaming in both sides, the response can be sended from both sides whenever they intend to
  ```
## Using the api
- .proto file generate both client and server side code, the apis are called in the client side and then implemented in the server side
- Server side, the server implements the methods declared by the service and runs a gRPC server to handle the client calls. The gRPC infrastructure decodes incoming requests, execute services methods and encode service responses
- Client, has a local object known as stub that implements the same methods as the service. The client can then just call those methods on the local object. After that, all the parameters get stored in the correct buffer message type, after that it sends it to the server and the server answer the request back

## Synchronous vs Asynchronous
- Synchronous RPC calls can block until the response get received. On the other hand beeing Async or Sync may depend of the use case and the go language may offer this operations out of the box   

## RPC Life Cycle