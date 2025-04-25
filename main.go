package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "testGRPC/gen/proto"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (u *userServiceServer) GetUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Получен запрос на пользователя с ID", request.Id)

	return &pb.UserResponse{
		Id:    request.Id,
		Name:  "Иван Иванов",
		Email: "hello@gmail.com",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка запуска сервера", err)
	}

	grpcServer := grpc.NewServer()
	log.Println("GRPS сервер запущен на порту 50051")

	pb.RegisterUserServiceServer(grpcServer, &userServiceServer{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка запуска сервера", err)
	}
}
