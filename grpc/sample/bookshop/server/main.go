package main

import (
	"context"
	"log"
	"net"

	pb "bookshop/server/pb/inventory"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func getSampleBooks() []*pb.Book {

	return []*pb.Book{
		{Title: "You can win", Author: "Unknown", PageCount: 121, Language: "english"},
		{Title: "You can win 2", Author: "Unknown 2", PageCount: 121, Language: "english"},
		{Title: "You can win 3", Author: "Unknown 3", PageCount: 121, Language: "english"},
	}
}

func main() {

	ls, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen tcp on port:9000 %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)

	pb.RegisterInventoryServer(s, &server{})
	err = s.Serve(ls)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
