package main

import (
	"context"
	"fmt"
	"fulfillment/database"
	pb "fulfillment/fulfillment"
	"log"
	"net"

	"fulfillment/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	pb.FulFillmentServer
	DB *gorm.DB
}


func main() {
	db := database.Connection()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFulFillmentServer(grpcServer, &Server{DB : db})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
	log.Println("Server started running on port 9000")
}

func(s *Server) RegisterDeliveryAgent(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req.Username == "" || req.Password == "" || req.City == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing user data")
	}
	_, err := database.FindByUsername(s.DB, req.Username)
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "User already exists")
	}
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		City:  req.City,
		Availability: models.AVAILABLE,
	}

	err = s.DB.Create(&user).Error

	if err != nil {
		errorString := fmt.Sprintf("error storing the user: %v", err)
		return nil, status.Errorf(codes.Unknown, errorString)
	}

	response := &pb.RegisterResponse{
		Message: "Successfully Registered!",
	}

	return response, nil
}

