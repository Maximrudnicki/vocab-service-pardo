package main

import (
	"context"

	"vocab_service/cmd/model"
	u "vocab_service/cmd/utils"
	pb "vocab_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetWords, CreateWord, DeleteWord, UpdateWord

func (s *Server) GetWords(in *pb.VocabRequest, stream pb.VocabService_GetWordsServer) error {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return err
	}

	words, words_err := s.WordRepository.FindByUserId(userId)
	if words_err != nil {
		return words_err
	}

	for _, word := range words {
		stream_err := stream.Send(&pb.VocabResponse{
			Id:              word.Id,
			Word:            word.Word,
			Definition:      word.Definition,
			CreatedAt:       u.ToTimestamp(word.CreatedAt),
			IsLearned:       word.IsLearned,
			Cards:           word.Cards,
			WordTranslation: word.WordTranslation,
			Constructor:     word.Constructor,
			WordAudio:       word.WordAudio,
		})
		if stream_err != nil {
			return stream_err
		}
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

	err = s.WordRepository.Save(newWord)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteWord(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, err
	}

	if isOwner := s.WordRepository.IsOwnerOfWord(userId, in.WordId); isOwner {
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
		Id:         in.Id,
		Definition: in.Definition,
		UserId:     userId,
	}

	if isOwner := s.WordRepository.IsOwnerOfWord(userId, in.Id); isOwner {
		err = s.WordRepository.Update(updatedWord)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"You are not allowed to update the word",
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ManageTrainings(ctx context.Context, in *pb.ManageTrainingsRequest) (*emptypb.Empty, error) {
	userId, err := u.GetUserIdFromToken(in.Token)
	if err != nil {
		return nil, err
	}

	if isOwner := s.WordRepository.IsOwnerOfWord(userId, in.Id); !isOwner {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"You are not allowed to manage trainings for this word",
		)
	}

	err_mt := s.WordRepository.ManageTrainings(in.Res, in.Training, in.Id)
	if err_mt != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Please check the training name and other arguments",
		)
	}

	return &emptypb.Empty{}, nil
}
