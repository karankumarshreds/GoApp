package service

import (
	"github.com/karankumarshreds/GoApp/internal/core"
	"github.com/karankumarshreds/GoApp/internal/dto"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
	ports "github.com/karankumarshreds/GoApp/internal/ports/repository"
)

type AccountService struct {
	repo ports.AccountRepositoryPort
}

func NewAccountService( repo ports.AccountRepositoryPort ) AccountService{
	return AccountService{repo}
}

func (as AccountService) CreateAccount(account dto.AccountCreateDto) (*dto.AccountCreateResponse, *custom_errors.CustomError) {
	a := core.Account{
		CustomerId:  account.CustomerId,
		AccountType: account.AccountType,
		Amount:      account.Amount,
	}
	newAccount, err := as.repo.Save(a)
	if  err != nil {
		return nil, err
	}
	return &dto.AccountCreateResponse{newAccount.AccountId}, nil
}