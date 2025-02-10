package main

import (
	"context"
	"log"
	"time"

	pb "github.com/yourusername/e-learning/studentpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetStudentRequest{Id: "12345"}
	res, err := client.GetStudent(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Printf("Student: %s (%s)", res.Student.Name, res.Student.Email)
}
