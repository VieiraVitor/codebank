package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/vieiravitor/codebank/dto"
	"github.com/vieiravitor/codebank/infrastructure/grpc/pb"
	"github.com/vieiravitor/codebank/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionService struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
	pb.UnimplementedPaymentServiceServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) Payment(ctx context.Context, paymentRequest *pb.PaymentRequest) (*empty.Empty, error) {
	transactionPayload := dto.TransactionPayload{
		Name:            paymentRequest.GetCreditCard().Name,
		Number:          paymentRequest.CreditCard.GetNumber(),
		ExpirationMonth: paymentRequest.CreditCard.GetExpirationMonth(),
		ExpirationYear:  paymentRequest.CreditCard.GetExpirationYear(),
		CVV:             paymentRequest.CreditCard.GetCvv(),
		Amount:          paymentRequest.GetAmount(),
		Store:           paymentRequest.GetStore(),
		Description:     paymentRequest.GetDescription(),
	}

	transaction, err := t.ProcessTransactionUseCase.ProcessTransaction(transactionPayload)
	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}
	if transaction.Status != "approved" {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, "transaction rejected by the bank")
	}
	return &empty.Empty{}, nil
}
