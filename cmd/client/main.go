package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/JokerTrickster/grpc_go/pkg/api/v1"
	game "github.com/JokerTrickster/grpc_go/pkg/game"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// 서버에 연결 (grpc.Dial)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Greeter와 Game 서비스 클라이언트 생성
	greeterClient := pb.NewGreeterClient(conn)
	gameClient := game.NewGameClient(conn)

	// 기본 이름 설정(명령행 인자 사용 가능)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// 짧은 타임아웃을 갖는 컨텍스트 생성
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Greeter 서비스의 SayHello 호출
	helloResp, err := greeterClient.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", helloResp.GetMessage())

	// Greeter 서비스의 SayHelloAgain 호출
	helloAgainResp, err := greeterClient.SayHelloAgain(ctx, &pb.HelloRequest{Name: "test"})
	if err != nil {
		log.Fatalf("could not greet again: %v", err)
	}
	log.Printf("Greeting: %s", helloAgainResp.GetMessage())

	// 게임 데이터 설정을 위한 컨텍스트 생성 (타임아웃 2초)
	ctxGame, cancelGame := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelGame()

	// Game 서비스의 SetGameData 호출
	setReq := &game.RequestGameInfo{
		Round:  1,
		Map:    []int32{1, 2, 3},
		RoomID: 123,
	}
	// SetGameData는 google.protobuf.Empty를 반환함
	_, err = gameClient.SetGameData(ctxGame, setReq)
	if err != nil {
		log.Fatalf("could not set game data: %v", err)
	}
	log.Println("SetGameData succeeded")

	// Game 서비스의 GetGameData 호출
	getResp, err := gameClient.GetGameData(ctxGame, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not get game data: %v", err)
	}
	log.Printf("Game Data: Round=%d, Map=%v, RoomID=%d", getResp.Round, getResp.Map, getResp.RoomID)
}
