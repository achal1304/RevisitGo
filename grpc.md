Sure! Here's the updated documentation with the added comparison between **HTTP** and **RPC**.

---

### RPC (Remote Procedure Call) Interview Prep Documentation

#### **What is RPC?**

Remote Procedure Call (RPC) is a protocol used for requesting services from a program located on another computer (commonly in a distributed system). It allows a client to call functions or procedures in a remote server as if they were local function calls.

#### **Key Concepts**:

1. **Client-Server Model**: RPC follows a client-server model, where the client sends a request to a server and the server responds with the requested data.
2. **Stub**: The client-side representation of the server function, which acts as an intermediary. It hides the complexity of communication over the network.
3. **Transport Layer**: The communication layer (usually TCP, HTTP) used for transferring data between client and server.
4. **Serialization/Deserialization**: The process of converting data structures into a byte stream for transmission (serialization) and converting them back into data structures on the receiving end (deserialization).
5. **Procedure Signature**: Defines the functions that can be called over the RPC. It includes the function name, parameters, and return types.

#### **Types of RPC**:

1. **Synchronous RPC**: The client waits for the server to complete the requested task and sends a response back before proceeding.
2. **Asynchronous RPC**: The client does not wait for the server’s response and can continue processing. The response is handled via callbacks or event handlers.

#### **Popular RPC Implementations**:

* **gRPC (Google RPC)**: A high-performance RPC framework that uses HTTP/2 for communication, Protocol Buffers (protobuf) for serialization, and supports many programming languages. It’s commonly used in microservices architectures.
* **JSON-RPC**: A lightweight RPC protocol that uses JSON for encoding messages. It is transport-agnostic (works over TCP, HTTP, WebSocket, etc.).
* **XML-RPC**: An RPC protocol that uses XML to encode the messages and HTTP as the transport protocol.

---

### **RPC vs HTTP: Key Differences**

| **Feature**         | **RPC**                                                                                                 | **HTTP**                                                                                                                       |
| ------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| **Purpose**         | Allows clients to call functions on a remote server as if they were local.                              | Transfers data between a client and a server over the web.                                                                     |
| **Protocol**        | RPC can work over any transport protocol, such as TCP, HTTP, etc.                                       | HTTP is a specific protocol used to transfer hypertext.                                                                        |
| **Communication**   | Typically uses binary or specialized encoding formats like Protocol Buffers (Protobuf).                 | Uses text-based formats like JSON or XML for data transfer.                                                                    |
| **Performance**     | Faster and more efficient (especially when using binary formats like Protobuf).                         | Slower due to textual formats and additional overhead (headers, etc.).                                                         |
| **Connection Type** | Supports both synchronous and asynchronous communication.                                               | Primarily synchronous communication (though HTTP/2 supports multiplexing).                                                     |
| **Use Case**        | Suitable for tightly coupled systems, like microservices or distributed systems.                        | Used for web browsing, RESTful APIs, and stateless communication.                                                              |
| **State**           | RPC calls are typically stateful, meaning the state of the application can be maintained between calls. | HTTP is stateless, meaning each request is independent, and the server doesn’t store the state of the client between requests. |
| **Ease of Use**     | More complex to implement and set up due to custom protocols and service definitions.                   | Simpler to implement using standard web technologies and libraries.                                                            |

---

### **gRPC Overview**:

1. **Protocol Buffers (Protobuf)**: The default interface definition language (IDL) used by gRPC for defining service methods and message types. Protobuf offers efficient serialization and is language-agnostic.

2. **Service Definition**: gRPC services are defined in `.proto` files.

   * Example service definition in a `.proto` file:

   ```proto
   syntax = "proto3";

   service Greeter {
       rpc SayHello (HelloRequest) returns (HelloReply);
   }

   message HelloRequest {
       string name = 1;
   }

   message HelloReply {
       string message = 1;
   }
   ```

3. **gRPC Methods**:

   * **Unary RPC**: A single request from the client and a single response from the server.
   * **Server Streaming RPC**: The client sends a single request, and the server sends a stream of responses.
   * **Client Streaming RPC**: The client sends a stream of requests to the server, and the server sends a single response.
   * **Bidirectional Streaming RPC**: Both the client and server send streams of messages to each other.

4. **Useful Commands**:

   * **Compile .proto file**: To generate Go or other language-specific code from a `.proto` file:

     ```bash
     protoc --go_out=. --go-grpc_out=. service.proto
     ```
   * **Start a gRPC server** (Go):

     ```go
     server := grpc.NewServer()
     pb.RegisterGreeterServer(server, &GreeterServer{})
     lis, err := net.Listen("tcp", ":50051")
     if err != nil {
         log.Fatalf("failed to listen: %v", err)
     }
     if err := server.Serve(lis); err != nil {
         log.Fatalf("failed to serve: %v", err)
     }
     ```

5. **Client Side**:

   * **Calling an RPC** (Go example):

   ```go
   conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
   if err != nil {
       log.Fatalf("did not connect: %v", err)
   }
   defer conn.Close()
   c := pb.NewGreeterClient(conn)
   r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
   if err != nil {
       log.Fatalf("could not greet: %v", err)
   }
   fmt.Println(r.GetMessage())
   ```

---

### **Key gRPC Concepts**:

1. **Service**: A service is a collection of methods defined in `.proto` files.
2. **Method**: A single function exposed over RPC.
3. **Messages**: Input and output data for the service methods, defined as messages in `.proto`.
4. **Streaming**: gRPC supports server, client, and bidirectional streaming, enabling efficient communication for large datasets.

#### **gRPC Performance Advantages**:

1. **HTTP/2**: gRPC uses HTTP/2, which supports multiplexed streams, connection reuse, and better performance over HTTP/1.x.
2. **Protocol Buffers**: Protobuf is more compact and faster than JSON and XML.
3. **Bidirectional Streaming**: Allows efficient real-time communication between client and server.
4. **Language Agnostic**: gRPC supports multiple languages, including Go, Java, Python, C++, etc.

---

### **Common gRPC Keywords**:

* **`service`**: Defines a service with a set of RPC methods.
* **`rpc`**: Defines a method in the service.
* **`message`**: Defines a message type that will be used as a parameter or return type in RPC methods.
* **`option`**: Used to configure the behavior of services and methods, such as specifying the language to generate code for.

#### **Use Case for RPC**:

* **Microservices**: gRPC is a common choice for microservices communication due to its speed and support for multiple programming languages.
* **Real-Time Systems**: For use cases requiring real-time communication, gRPC's support for bidirectional streaming makes it ideal.

### **Key gRPC Commands**:

1. **Compile Protobuf**:
   To generate Go code from `.proto` files:

   ```bash
   protoc --go_out=. --go-grpc_out=. service.proto
   ```
2. **Start gRPC Server**:
   Create and start a gRPC server in Go:

   ```go
   server := grpc.NewServer()
   pb.RegisterServiceServer(server, &ServiceServer{})
   lis, err := net.Listen("tcp", ":50051")
   if err != nil {
       log.Fatalf("failed to listen: %v", err)
   }
   if err := server.Serve(lis); err != nil {
       log.Fatalf("failed to serve: %v", err)
   }
   ```
3. **gRPC Client Call**:
   Call the gRPC server from a client:

   ```go
   conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
   if err != nil {
       log.Fatalf("did not connect: %v", err)
   }
   defer conn.Close()
   c := pb.NewServiceClient(conn)
   response, err := c.MethodName(context.Background(), &pb.RequestType{})
   ```

---

This comprehensive documentation should help you prepare for an interview on RPC, gRPC, and the differences between HTTP and RPC. You can use it as a quick reference for understanding the key concepts, common commands, and their use cases. Let me know if you'd like more details or examples!
