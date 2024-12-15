package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/myeunee/FSSN_gRPC/lec-07-prg-01-hello_gRPC"
	"google.golang.org/grpc"
)

func main() {
	// (3) gRPC 통신 채널 생성
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// (4) 통신 채널로 stub 생성
	client := pb.NewMyServiceClient(conn)

	// (5) 원격 함수에 전달할 메시지 만듬
	request := &pb.MyNumber{Value: 4}

	// (6) gRPC 원격 함수 호출
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.MyFunction(ctx, request)
	if err != nil {
		log.Fatalf("Failed to call MyFunction: %v", err)
	}

	// (7) 결과 출력
	fmt.Printf("gRPC result: %d\n", response.Value)
}
