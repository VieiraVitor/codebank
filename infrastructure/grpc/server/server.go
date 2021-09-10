package server

import (
	"log"
	"net"

	"github.com/vieiravitor/codebank/infrastructure/grpc/pb"
	"github.com/vieiravitor/codebank/infrastructure/grpc/service"
	"github.com/vieiravitor/codebank/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	listen, err := net.Listen("tpc", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("could not listen to tpc port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	grpcServer.Serve(listen)
}
