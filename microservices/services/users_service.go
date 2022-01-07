package services

import (
	"fmt"

	"github.com/Nagateja1034845/microservices/domain/employees"
	"github.com/Nagateja1034845/microservices/domain/httperrors"
)

var (
	EmpsService = empsService{}

	registeredEmps       = map[int64]*employees.Employee{}
	currentEmpId   int64 = 1
)

type empsService struct{}

func (service empsService) Create(emp employees.Employee) (*employees.Employee, *httperrors.HttpError) {
	if err := emp.Validate(); err != nil {
		return nil, err
	}

	emp.EmpId = currentEmpId
	currentEmpId++

	registeredEmps[emp.EmpId] = &emp

	return &emp, nil
}

func (service empsService) Get(EmpId int64) (*employees.Employee, *httperrors.HttpError) {
	if user := registeredEmps[EmpId]; user != nil {
		return user, nil
	}
	return nil, httperrors.NewNotFoundError(fmt.Sprintf("user %d not found", EmpId))
}

/*func (service empsService) Update(emp employees.Employee, isPartial bool) (*employees.Employee, *httperrors.HttpError) {
	current, err:=UpdateEmp(emp.EmpId)
	if err!=nil{
		return nil,err
	}
	if isPartial{
		if emp.EmpName!=""{
			current.EmpName=emp.EmpName
		}
		if emp.EmpAge != 0{
current.EmpAge=emp.EmpAge
		}
		if emp.EmpEmail != ""{
			current.EmpEmail=emp.EmpEmail

		}

	}else{

		current.EmpName=emp.EmpName
		current.EmpAge=emp.EmpAge
		current.EmpEmail=emp.EmpEmail
	}
	if err:=current.Update();err!=nil{
		return nil,err
	}
	return current,nil
}*/
