package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Hashimoto-Noriaki/e-learning/studentpb"
	"google.golang.org/grpc"
)

// StudentServiceServer の実装
type server struct {
	pb.UnimplementedStudentServiceServer
}

// GetStudent の実装
func (s *server) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {
	student := &pb.Student{
		Id:    req.GetId(),
		Name:  "Taro Yamada",
		Email: "taro@example.com",
	}

	return &pb.GetStudentResponse{Student: student}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStudentServiceServer(grpcServer, &server{})

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
