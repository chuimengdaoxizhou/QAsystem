package text_generate_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "qasystem/text-generate-client/proto"
	"time"
)

func GenerateToken(prompt string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接服务器失败,%s", err)
		return
	}
	defer conn.Close()

	client := pb.NewModelServiceClient(conn)

	req := &pb.GenerationRequest{
		Prompt: prompt,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// 注意这里的ctx类型是正确的context.Context，不是gin.Context
	stream, err := client.GenerateContentStream(ctx, req)
	if err != nil {
		log.Fatalf("调用流式RPC失败: %v", err)
	}

	fmt.Println("AI Response:")
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("接收流数据出错: %v", err)
		}
		fmt.Print(resp.Token)
	}
	fmt.Println("\n------")
}
