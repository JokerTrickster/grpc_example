package main

import (
	"context"
	"log"
	"net"

	v1 "github.com/JokerTrickster/grpc_go/pkg/api/v1"
	game "github.com/JokerTrickster/grpc_go/pkg/game"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = ":50051"
)

// server는 Greeter와 Game 두 서비스를 모두 구현합니다.
type server struct {
	v1.UnimplementedGreeterServer
	game.UnimplementedGameServer
}

// Greeter 서비스의 SayHello 메서드 구현
func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	log.Printf("Received SayHello: %v", in.GetName())
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// Greeter 서비스의 SayHelloAgain 메서드 구현
func (s *server) SayHelloAgain(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	log.Printf("Received SayHelloAgain: %v", in.GetName())
	return &v1.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

// Game 서비스의 SetGameData 메서드 구현
func (s *server) SetGameData(ctx context.Context, in *game.RequestGameInfo) (*emptypb.Empty, error) {
	log.Printf("Received SetGameData: Round=%d, Map=%v, RoomID=%d", in.Round, in.Map, in.RoomID)
	// 여기서 입력된 게임 데이터를 처리합니다.
	return &emptypb.Empty{}, nil
}

// Game 서비스의 GetGameData 메서드 구현
func (s *server) GetGameData(ctx context.Context, in *emptypb.Empty) (*game.ResponseGameInfo, error) {
	log.Printf("Received GetGameData request")
	// 예시로, 기본 값을 반환합니다.
	resp := &game.ResponseGameInfo{
		Round:  1,
		Map:    []int32{1, 2, 3},
		RoomID: 123,
	}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// Greeter 서비스 등록
	v1.RegisterGreeterServer(s, &server{})
	// Game 서비스 등록
	game.RegisterGameServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
