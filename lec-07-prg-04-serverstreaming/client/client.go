package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/myeunee/FSSN_gRPC/lec-07-prg-04-serverstreaming"
	"google.golang.org/grpc"
)

func recvMessage(client pb.ServerStreamingClient) {
	ctx := context.Background()
	req := &pb.Number{Value: 5}
	stream, err := client.GetServerResponse(ctx, req)
	if err != nil {
		log.Fatalf("[에러] Failed to call GetServerResponse: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("[에러] Failed to receive message: %v", err)
		}
		fmt.Printf("[server to client] %s\n", msg.GetMessage())
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[에러] Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewServerStreamingClient(conn)
	recvMessage(client)
}
