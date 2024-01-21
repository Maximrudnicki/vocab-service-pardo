package main

import (
	"log"
	"net"
	"vocab_service/cmd/config"
	"vocab_service/cmd/model"
	"vocab_service/cmd/repository"
	pb "vocab_service/proto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.VocabServiceServer
	WordRepository repository.WordRepository
}

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)

	db_table_err := db.Table("words").AutoMigrate(&model.Word{})
	if db_table_err != nil {
		log.Fatalf("Databese table error: %v\n", db_table_err)
	}

	//Init Repository
	wordRepository := repository.NewWordRepositoryImpl(db)

	// Start GRPC Server
	lis, err := net.Listen("tcp", loadConfig.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", loadConfig.GRPCPort)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	pb.RegisterVocabServiceServer(s, &Server{WordRepository: wordRepository})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
