package ports

import (
	"github.com/karankumarshreds/GoApp/internal/core"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
)

type AccountRepositoryPort interface {
	Save(account core.Account) (*core.Account, *custom_errors.CustomError)
}
