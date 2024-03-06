package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"fulfillment/database"
	"fulfillment/fulfillment"
	pb "fulfillment/fulfillment"
	"log"
	"net"
	"strings"

	"fulfillment/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
		errorString := fmt.Sprintf("Error storing the user: %v", err)
		return nil, status.Errorf(codes.Unknown, errorString)
	}

	response := &pb.RegisterResponse{
		Message: "Successfully Registered!",
	}

	return response, nil
}


func(s *Server) AssignDeliveryAgent(ctx context.Context, req *pb.AssignRequest) (*pb.AssignResponse, error) {
	delivery_agent, err := database.FindAvailableDeliveryAgentByCity(s.DB, req.City)

	if err != nil {
		errorString := fmt.Sprintf("No delivery agent available: %v", err)
		return nil, status.Error(codes.Unavailable, errorString)
	}

	delivery := &models.Delivery{
		OrderID: uint64(req.OrderId),
		City: req.City,
		DeliveryAgent: delivery_agent,
	}

	err = s.DB.Create(&delivery).Error

	if err != nil {
		errorString := fmt.Sprintf("Error storing the delivery: %v", err)
		return nil, status.Errorf(codes.Unknown, errorString)
	}

	response := &pb.AssignResponse{
		Message: "Successfully Assigned a delivery agent!",
	}

	return response, nil
}

func (s *Server) UpdateDeliveryAgentAvailability(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	user, err := s.getCredentials(ctx)
	if err != nil {
		errorString := fmt.Sprintf("%v", err)
		return nil, status.Error(codes.Unauthenticated, errorString)
	}
	user.Availability = models.AVAILABLE
	s.DB.Save(&user)

	

	response := &pb.UpdateResponse{
		Message: "Order delivered and the delivery agent is now available",
	}
	return response, nil
}

func (s *Server) FetchAllDeliveriesForAnAgent(ctx context.Context, req *pb.FetchDeliveriesRequest) (*pb.FetchDeliveriesResponse, error) {
	user, err := s.getCredentials(ctx)
	if err != nil {
		errorString := fmt.Sprintf("%v", err)
		return nil, status.Error(codes.Unauthenticated, errorString)
	}
	deliveries := database.FetchDeliveriesForAnAgent(s.DB, user.ID)
	response := &pb.FetchDeliveriesResponse{}
	for _, delivery := range deliveries {
		deliveryResponse := &fulfillment.Delivery{Id: delivery.ID,  OrderId: delivery.OrderID, City: delivery.City}
		response.Deliveries = append(response.Deliveries, deliveryResponse)
	}
	return response, nil
}

func(s *Server) getCredentials(ctx context.Context) (models.User, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	authHeader, ok := md["authorization"]

	if !ok || len(authHeader) == 0 {
		return models.User{}, errors.New("authorization header not found")
	}

	authParts := strings.Fields(authHeader[0])
	if len(authParts) != 2 || authParts[0] != "Basic" {
		return models.User{}, errors.New("invalid Authorization header format")
	}

	decodedCredentials, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		return models.User{}, errors.New("error decoding base64 credentials")
	}

	credentials := strings.SplitN(string(decodedCredentials), ":", 2)
	if len(credentials) != 2 {
		return models.User{}, errors.New("invalid credentials format")
	}

	username := credentials[0]
	password := credentials[1]

	var user models.User
	err = s.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return models.User{}, errors.New("Invalid credentials")
	}

	return user, nil
}
