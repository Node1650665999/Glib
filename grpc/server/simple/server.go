package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/node1650665999/Glib/grpc/proto"
)

//SearchService 定义服务，需实现了SearchServiceServer，这样该服务才能注册
type SearchService struct{}
//Search 实现了SearchServiceServer接口定义的方法,实现该方法就能填充我们的业务逻辑
func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	//创建 gRPC Server,用来注册服务
	server := grpc.NewServer()
	//注册服务
	pb.RegisterSearchServiceServer(server, &SearchService{})
	//监听tcp请求
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	//处理
	server.Serve(lis)
}
