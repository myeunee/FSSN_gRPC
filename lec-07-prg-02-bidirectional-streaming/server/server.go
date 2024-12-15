package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/myeunee/FSSN_gRPC/lec-07-prg-02-bidirectional-streaming"
	"google.golang.org/grpc"
)

type BidirectionalService struct {
	pb.UnimplementedBidirectionalServer
}

func (s *BidirectionalService) GetServerResponse(stream pb.Bidirectional_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC bidirectional streaming.")
	for {
		message, err := stream.Recv() // 클라이언트로부터 메시지 수신
		if err != nil {
			return err
		}
		fmt.Printf("[client to server] %s\n", message.GetMessage())
		response := &pb.Message{Message: message.GetMessage()} // 수신 메시지를 리턴
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterBidirectionalServer(server, &BidirectionalService{})
	fmt.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
