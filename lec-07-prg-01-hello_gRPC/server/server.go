package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/myeunee/FSSN_gRPC/lec-07-prg-01-hello_gRPC"
	"google.golang.org/grpc"
)

// (4) MyServiceImpl 구조체를 정의 -> MyServiceServer 인터페이스 구현
type MyServiceImpl struct {
	pb.UnimplementedMyServiceServer
}

// (5) proto 파일의 rpc 함수에 대응하는 MyFunction 구현
func (s *MyServiceImpl) MyFunction(ctx context.Context, req *pb.MyNumber) (*pb.MyNumber, error) {
	// (5.2) 응답 메시지 생성
	response := &pb.MyNumber{
		Value: pb.MyFunc(req.Value),
	}
	return response, nil
}

func main() {
	// (6) gRPC 서버 생성
	server := grpc.NewServer()

	// (7) MyServiceImpl을 서버에 등록
	pb.RegisterMyServiceServer(server, &MyServiceImpl{})

	// (8) 포트 열기 + 서버 실행
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	fmt.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
