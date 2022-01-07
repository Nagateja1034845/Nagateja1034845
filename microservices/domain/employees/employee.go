package employees

import (
	"github.com/Nagateja1034845/microservices/domain/httperrors"
)

type Employee struct {
	EmpId    int64  `json:"EmpId"`
	EmpName  string `json:"EmpName"`
	EmpAge   int64  `json:"EmpAge"`
	EmpEmail string `json:"EmpEmail"`
}

func (emp Employee) Validate() *httperrors.HttpError {
	if emp.EmpName == "" {
		return httperrors.NewBadRequestError("invalid employee name")
	}
	if emp.EmpEmail == "" {
		return httperrors.NewBadRequestError("invalid email address")
	}
	return nil
}
