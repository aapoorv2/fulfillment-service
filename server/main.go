package main

import (
	"gorm.io/gorm"
	pb "fulfillment/fulfillment"

)

type DeliveryServer struct {
	DB *gorm.DB
	pb.FulFillmentServer
}

type UserServer struct {
	DB *gorm.DB
	pb.FulFillmentServer
}