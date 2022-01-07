package controllers

import (
	"net/http"
	"strconv"

	"github.com/Nagateja1034845/microservices/domain/employees"
	"github.com/Nagateja1034845/microservices/domain/httperrors"
	"github.com/Nagateja1034845/microservices/services"
	"github.com/gin-gonic/gin"
)

var (
	EmpsController = empsController{}
)

type empsController struct{}

func respond(c *gin.Context, isXml bool, httpCode int, body interface{}) {
	if isXml {
		c.XML(httpCode, body)
		return
	}
	c.JSON(httpCode, body)
}

func (controller empsController) Create(c *gin.Context) {
	var emp employees.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid json body")
		c.JSON(httpErr.Code, httpErr)
		return
	}
	createdUser, err := services.EmpsService.Create(emp)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	// return created user
	c.JSON(http.StatusCreated, createdUser)
}

func (controller empsController) Get(c *gin.Context) {
	isXml := c.GetHeader("Accept") == "application/xml"

	empId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httpErr := httperrors.NewBadRequestError("invalid user id")
		respond(c, isXml, httpErr.Code, httpErr)
		return
	}

	emp, getErr := services.EmpsService.Get(empId)
	if getErr != nil {
		respond(c, isXml, getErr.Code, getErr)
		return
	}
	respond(c, isXml, http.StatusOK, emp)
}

/*func Update(c *gin.Context){
empId, empErr:=strconv.ParseInt(c.Param("id"),10,64)
if empErr!=nil{
	httpErr := httperrors.NewBadRequestError("invalid user id")
	c.JSON(httpErr.Code,httpErr)
}
var emp employees.Employee
if err := c.ShouldBindJSON(&emp); err != nil {
	httpErr := httperrors.NewBadRequestError("invalid json body")
	c.JSON(httpErr.Code, httpErr)
	return
}
emp.EmpId=empId
isPartial:=c.Request.Method==http.MethodPatch

result, err:= services.Updateemp(emp,isPartial)
if err!=nil{
	c.JSON(err.Code,err)
	return
}
c.JSON(http.StatusOK,result)
}*/
