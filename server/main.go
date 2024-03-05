package main

import (
	"fulfillment/database"
	pb "fulfillment/fulfillment"
	"log"
	"net"

	"google.golang.org/grpc"
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