# Go 프로그램 빌드를 위한 Makefile

# 설정
PROTO_DIR := pkg/proto
GO_OUT := pkg/proto  # 생성된 코드를 같은 폴더에 넣거나 다른 경로로 지정 가능
GO_PROTO_PACKAGE := pkg/proto


# Protocol Buffers 정의 파일 경로
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Protocol Buffers를 사용하여 Go 코드 생성
generate:
	protoc -I=$(PROTO_DIR) --go_out=$(GO_OUT) --go-grpc_out=$(GO_OUT) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_FILES)

# 빌드
build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/client cmd/client/main.go

# 실행
run-server:
	bin/server

run-client:
	bin/client