package utils

import (
	pb "vocab_service/proto"
	"context"
	"log"
)

func ValidateToken(c pb.AuthenticationServiceClient, token string) (*pb.UserIdResponse, error) {
	log.Println("---Validate Token was invoked---")

	req := &pb.TokenRequest{
		Token: token,
	}	
	
	res, err := c.GetUserId(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while getting ID: %v\n", err)
		return nil, err
	}

	log.Printf("ID: %v\n", res)
	return res, nil
}
