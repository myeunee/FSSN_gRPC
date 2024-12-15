package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/myeunee/FSSN_gRPC/lec-07-prg-02-bidirectional-streaming"
	"google.golang.org/grpc"
)

func generateMessages() <-chan *pb.Message {
	messages := []string{
		"message #1",
		"message #2",
		"message #3",
		"message #4",
		"message #5",
	}
	ch := make(chan *pb.Message)
	go func() {
		for _, msg := range messages {
			fmt.Printf("[client to server] %s\n", msg)
			ch <- &pb.Message{Message: msg}
		}
		close(ch)
	}()
	return ch
}

func sendMessage(client pb.BidirectionalClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.GetServerResponse(ctx)
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}

	for msg := range generateMessages() {
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive response: %v", err)
		}
		fmt.Printf("[server to client] %s\n", response.GetMessage())
	}
	stream.CloseSend()
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewBidirectionalClient(conn)
	sendMessage(client)
}
