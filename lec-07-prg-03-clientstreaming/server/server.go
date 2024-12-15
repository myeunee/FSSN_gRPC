package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/myeunee/FSSN_gRPC/lec-07-prg-03-clientstreaming"
	"google.golang.org/grpc"
)

type ClientStreamingService struct {
	pb.UnimplementedClientStreamingServer
}

// GetServerResponse handles client-streaming.
func (s *ClientStreamingService) GetServerResponse(stream pb.ClientStreaming_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC client-streaming.")
	count := 0
	for {
		_, err := stream.Recv() // 클라이언트로부터 메시지 수신
		if err != nil {
			if err.Error() == "EOF" {
				break // 스트리밍 종료
			}
			return err
		}
		count++
	}
	response := &pb.Number{Value: int32(count)}
	fmt.Printf("[server to client] %d\n", response.Value)
	return stream.SendAndClose(response) // 응답 전송 + 스트리밍 닫기
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterClientStreamingServer(server, &ClientStreamingService{})
	fmt.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
