# Introduction to gRPC
- It is a protocol that uses protobuf, to call a object from a GRPC Server, that by using protobuf enables the client to run it like it is a local object
- As another RPCS, it is based on defining a certain service using interfaces, mentioning the methods
- The server implements its interface and runs a GRPC server to handle the client calls
- On the client side, the client has a stub, that provides the same methods as the server
![Grpc representation](assets/grpc.png)
- Servers and clients can be writen in different programming languages
- gRPC uses [Protocol Buffers](Protocol-buffers/readme.md) by default, which is a different type of serialization of data such as JSON (we can still use json)

#  Quick Start
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
- The lifecycle is for types of RPCS that we mentioned before
### Unary RPC
1. - Client calls the stub method
   - Server is notified about the invocation with the clients metadata for this call, the method name and the deadline if disposed
----
2. - Server either send back its own initial metadata or wait a client request message, what happens first is app-specific
   - Metadata must be sent before any response 
----
3. - Once the server has the client request message it does whatever is necessary to populate the request with a response
   - The response is returned if sucessfull (together with status detail and optional trailling data)
---- 
4. - If status is OK, the client gets the response, which completes the call on the client Side
### Server Streamming RPC
- This one is almost the same but the difference is that the server sends all the messages and only on the final sends the server status details
### Client Streaming RPC
- Simmilar to the unary RPC, but the difference is that the client sends all the messages and gets one response back from the server with the server details
### Bidirectional Streaming RPC
- The client initiates the communications, invoking the methods and the server receiving the client metadata, method name and deadline. 
- The server can choose to send back its initial metadata or wait for the client to start streamming messages
- After all of that the proccess is app-specific, the 2 streams are independent
- Both can read and write messages in any order
- It works as a ping pong, or the server can wait the receival of messages, it is app-specific
### Deadlines/Timeout
- Clients can specify how much time they are willing to wait for an RPC to complete, before the RPC gets terminated with a DEADLINE_EXCEED error
### RPC Termination
- The conclusions of the call may not match between the two components (ex: The server may think the call went sucessfully, but on the client side it can have another kind of conclusion)
### Cancelling an RPC
- Both the server and the client can cancel an RPC at any time
- There are no rollbacks, what is done, is done
### Metadata
- It is information about a particular RPC call (such as authentication details)
- It is in a list of key-value pairs, there keys are Strings and values are typically strings, but can be also binary
- Keys are case insensitive and consist of ASCII letters, digits and special chars "-","_","."
- It cannot start with grpc-
- Binary valued keys end in "-bin", ASCII-valued keys do not
- Access to metadata is language dependent
- We can add custom key value pairs, which makes gRPC very flexible
### Channels
- gRPC channels are used to provide a connection in a specific host and port
- It is used when creating a client stub
- Clients can specify arguments to modify the behavior of the channel
- Channel has state, including connected and idle
# Basics Tutorial
- This example is a simple route mapping app
- Lets the clients know features on their route
- Create Summary of their route
- Exchange information such as traffic updates with the server and other clients
- Grpc becomes very usefull because it creates all the client and server and handles the communication himself. Also it has efficient serialization because of the buffers protocol, simple IDL and easy interface upgrating
## Notes about the tutorial
- The tutorial is here [Link](./basics-tutorial/)
- We defined the services and one message
- To generate the files we need to use this command:
  ```
    protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. route_guide.proto
  ```
- Note that we are saying that we want to generate it on the same directory we are. Also, we need to be on that exac directory as well because of the package that the generated files will retain
- After running this command it will generate two files: route_guide.pb.go, which stands for a protocol buffer to populate, serialize and retrieve request and response message types. route_guide_grpc.pb.go, which stands for a interface type (or stub) for clients to call with the methods defined in the RouteGuide service and a given interface type for servers to implement, also with the methods defined in the RouteGuide Service
- We will firstly implement the server using this interface