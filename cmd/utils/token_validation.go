package utils

import (
	pb "vocab_service/proto"
	"context"
	"log"
)

func ValidateToken(c pb.AuthenticationServiceClient, token string) (*pb.UserIdResponse, error) {
	req := &pb.TokenRequest{
		Token: token,
	}	
	
	res, err := c.GetUserId(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while getting ID: %v\n", err)
		return nil, err
	}

	return res, nil
}
