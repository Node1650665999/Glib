package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	pb "github.com/node1650665999/Glib/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"net"
	"runtime/debug"
)

//SearchService 定义服务，需实现了SearchServiceServer，这样该服务才能注册
type SearchService struct{}

//Search 实现了SearchServiceServer接口定义的方法,实现该方法就能填充我们的业务逻辑
func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {

	c    := GetTLSCredentials()
	opts := []grpc.ServerOption{
		//证书
		grpc.Creds(c),
		//注册中间件
		grpc_middleware.WithUnaryServerChain(
			RecoveryInterceptor,
			LoggingInterceptor,
		),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}

func GetTLSCredentials() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("../../cert/server.pem", "../../cert/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err  := ioutil.ReadFile("../../cert/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	return c
}


type ErrorData struct {
	Code int
	Msg  string
}

func ErrorDataString(code int, msg string) string  {
	str,_ := json.Marshal(ErrorData{code, msg})
	return string(str)
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//模拟超时
	//time.Sleep(10 * time.Second)

	//模拟错误
	/*if req!= "gRPC" {
		return nil, status.Errorf(codes.Internal, "cannot access!!!")
	}*/

	//模拟异常,将触发RecoveryInterceptor中间件
	//panic("some thing error")

	//附带自定义业务错误
	/*if req!= "gRPC" {
		selfError := ErrorDataString(3000, "自定义业务参数")
		sts,_ := status.New(codes.Internal, "Rpc self define err").WithDetails(proto.MessageV1(selfError))
		return nil, sts.Err()
	}*/


	//前置操作
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	//后置操作
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}



func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//异常捕获
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
			status.New(codes.Internal, "Panic err").WithDetails(proto.MessageV1("asdsaf"))
		}
	}()

	return handler(ctx, req)
}
