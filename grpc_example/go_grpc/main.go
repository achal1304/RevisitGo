package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	example "github.com/revisitgo/grpc_example/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	example.UnimplementedMessagingServiceServer
}

// SendMessage handles the unary RPC
func (s *server) SendMessage(ctx context.Context, req *example.MessageRequest) (*example.MessageResponse, error) {
	log.Printf("Received message: %s", req.GetMessage())
	return &example.MessageResponse{Response: "Hello, " + req.GetMessage()}, nil
}

func (s *server) StreamMessages(req *example.StreamRequest, stream example.MessagingService_StreamMessagesServer) error {
	log.Printf("Streaming messages with prefix: %s", req.GetPrefix())

	for i := 0; i < 5; i++ {
		message := req.GetPrefix() + fmt.Sprintf(" - Message #%d", i)
		if err := stream.Send(&example.StreamResponse{Message: message}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil

}

func (s *server) StreamClientMessages(stream example.MessagingService_StreamClientMessagesServer) error {
	var fullMessage string
	for {
		// Try receiving the message from the client
		req, err := stream.Recv()
		if err == io.EOF {
			// EOF indicates the client has finished sending messages
			break
		}
		if err != nil {
			// Return error if something other than EOF occurs
			return fmt.Errorf("failed to receive message: %v", err)
		}
		// Append the received message to the fullMessage string
		fullMessage += req.GetMessage() + " "
	}
	fmt.Println("Received Message ", fullMessage)

	// Once the loop breaks (EOF), we send a response back to the client
	return stream.SendAndClose(&example.MessageResponse{Response: "Received: " + fullMessage})
}

func (s *server) BidirectionalStream(stream example.MessagingService_BidirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Received message: %s", req.GetMessage())

		// Respond back to the client
		if err := stream.Send(&example.MessageResponse{Response: "Acknowledged: " + req.GetMessage()}); err != nil {
			return err
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	example.RegisterMessagingServiceServer(s, &server{})

	reflection.Register(s)

	fmt.Println("Server is running on port 9000...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
