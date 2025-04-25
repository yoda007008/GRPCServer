package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "testGRPC/gen/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка подключения")
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetUser(ctx, &pb.UserRequest{Id: "123"})
	if err != nil {
		log.Fatalf("Получаем данные от пользователя ID: %s, Name: %s, Email: %s", res.Id, res.Name, res.Email)
	}
}
