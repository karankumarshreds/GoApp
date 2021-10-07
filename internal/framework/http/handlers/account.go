package handlers

import (
	"encoding/json"
	"github.com/karankumarshreds/GoApp/internal/dto"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
	"github.com/karankumarshreds/GoApp/internal/ports/service"
	"net/http"
)

type AccountHandlers struct {
	service ports.AccountServicePort
}

func NewAccountHandlers(service ports.AccountServicePort) AccountHandlers {
	return AccountHandlers{service}
}

func (ah *AccountHandlers) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var a dto.AccountCreateDto
	var err error
	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		custom_errors.NewBadRequestError("Bad request body").ErrorJsonResponse(w)
		return
	}
	res, err_ := ah.service.CreateAccount(a)
	if err_ != nil {
		err_.ErrorJsonResponse(w)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		custom_errors.NewInternalServerError("Unable to process the request")
		return
	}
}





