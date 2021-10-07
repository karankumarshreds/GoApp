package ports

import (
	"github.com/karankumarshreds/GoApp/internal/dto"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
)

type AccountServicePort interface {
	CreateAccount(account dto.AccountCreateDto) (*dto.AccountCreateResponse, *custom_errors.CustomError)
}

