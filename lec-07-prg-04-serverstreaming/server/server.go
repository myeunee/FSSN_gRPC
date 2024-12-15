package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/myeunee/FSSN_gRPC/lec-07-prg-04-serverstreaming"
	"google.golang.org/grpc"
)

func makeMessage(text string) *pb.Message {
	return &pb.Message{Message: text}
}

type ServerStreamingService struct {
	pb.UnimplementedServerStreamingServer
}

// GetServerResponse handles server-streaming.
func (s *ServerStreamingService) GetServerResponse(req *pb.Number, stream pb.ServerStreaming_GetServerResponseServer) error {
	messages := []*pb.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}
	fmt.Printf("Server processing gRPC server-streaming {%d}.\n", req.GetValue())
	for _, msg := range messages {
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("[에러] Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterServerStreamingServer(server, &ServerStreamingService{})
	fmt.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("[에러] Failed to serve: %v", err)
	}
}
