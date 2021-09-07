package usecase

import (
	"time"

	"github.com/vieiravitor/codebank/domain"
	"github.com/vieiravitor/codebank/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionPayload dto.TransactionPayload) (domain.Transaction, error) {
	creditCard := u.buildCreditCard(transactionPayload)
	ccBalanceAndLimit, err := u.TransactionRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}

	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance
	transaction := u.buildTransaction(transactionPayload, ccBalanceAndLimit)
	transaction.ProcessAndValidate(creditCard)
	err = u.TransactionRepository.SaveTransaction(*transaction, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	return *transaction, nil
}

func (u UseCaseTransaction) buildCreditCard(transactionPayload dto.TransactionPayload) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionPayload.Name
	creditCard.Number = transactionPayload.Number
	creditCard.ExpirationMonth = transactionPayload.ExpirationMonth
	creditCard.ExpirationYear = transactionPayload.ExpirationYear
	creditCard.CVV = transactionPayload.CVV
	return creditCard
}

func (u UseCaseTransaction) buildTransaction(transactionPayload dto.TransactionPayload, cc domain.CreditCard) *domain.Transaction {
	transaction := domain.NewTransaction()
	transaction.CreditCardId = cc.ID
	transaction.Amount = transactionPayload.Amount
	transaction.Store = transactionPayload.Store
	transaction.Description = transactionPayload.Description
	transaction.CreatedAt = time.Now()
	return transaction
}
