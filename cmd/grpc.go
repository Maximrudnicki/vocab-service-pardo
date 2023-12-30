package main

import (
	"context"
	"log"

	"vocab_service/cmd/model"
	u "vocab_service/cmd/utils"
	pb "vocab_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetWords, CreateWord, DeleteWord, UpdateWord

func (s *Server) GetWords(in *pb.VocabRequest, stream pb.VocabService_GetWordsServer) error {
	log.Println("-- GetWords Server --")
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return err
	}

	words, words_err := s.WordRepository.FindByUserId(userId)
	if words_err != nil {
		return words_err
	}

	for _, word := range words {
		stream.Send(&pb.VocabResponse{
			Id:         word.Id,
			Word:       word.Word,
			Definition: word.Definition,
		})
	}

	return nil
}

func (s *Server) CreateWord(ctx context.Context, in *pb.CreateRequest) (*emptypb.Empty, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, err
	}

	newWord := model.Word{
		Word:       in.Word,
		Definition: in.Definition,
		UserId:     userId,
	}

	s.WordRepository.Save(newWord)

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteWord(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, err
	}

	if is_owner := s.WordRepository.IsOwnerOfWord(userId, in.WordId); is_owner == true {
		s.WordRepository.Delete(in.WordId)
	} else {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"You are not allowed to delete the word",
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateWord(ctx context.Context, in *pb.UpdateRequest) (*emptypb.Empty, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, err
	}

	updatedWord := model.Word{
		Id: in.Id,
		Definition: in.Definition,
		UserId:     userId,
	}

	if is_owner := s.WordRepository.IsOwnerOfWord(userId, in.Id); is_owner == true {
		s.WordRepository.Update(updatedWord)
	} else {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"You are not allowed to update the word",
		)
	}

	return &emptypb.Empty{}, nil
}
