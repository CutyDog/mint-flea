package main

import (
	"log"
	"net"
	"os"

	accountv1 "github.com/CutyDog/mint-flea/proto/gen/account/v1"
	"github.com/CutyDog/mint-flea/services/account/internal/db"
	"github.com/CutyDog/mint-flea/services/account/internal/repo"
	"github.com/CutyDog/mint-flea/services/account/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db.ConnectDB()

	// Initialize repository
	accountRepo := repo.NewAccountRepository(db.DB)

	// Setup gRPC server
	addr := os.Getenv("GRPC_ADDR")
	s := grpc.NewServer()
	accountv1.RegisterAccountServiceServer(s, server.NewAccountServer(accountRepo))
	reflection.Register(s)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	log.Printf("account gRPC listening on %s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
