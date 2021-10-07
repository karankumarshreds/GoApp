package service

import (
	"github.com/karankumarshreds/GoApp/internal/core"
	"github.com/karankumarshreds/GoApp/internal/dto"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
	ports "github.com/karankumarshreds/GoApp/internal/ports/repository"
	"time"
)

type AccountService struct {
	repo ports.AccountRepositoryPort
}

func NewAccountService( repo ports.AccountRepositoryPort ) AccountService{
	return AccountService{repo}
}

func (as AccountService) CreateAccount(account dto.AccountCreateDto) (*dto.AccountCreateResponse, *custom_errors.CustomError) {
	a := core.Account{
		AccountId:   "",
		CustomerId:  account.CustomerId,
		OpeningDate: time.Now().Format(time.RFC3339),
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      "active",
	}
	newAccount, err := as.repo.Save(a)
	if  err != nil {
		return nil, err
	}
	return &dto.AccountCreateResponse{ AccountId: newAccount.AccountId}, nil
}