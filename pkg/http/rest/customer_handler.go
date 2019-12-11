package rest

import (
	"net/http"
	"template-api-go/pkg/customer"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type customerCtrl struct {
	svc customer.Service
}

func NewCustomerCtrl(svc customer.Service) *customerCtrl {
	return &customerCtrl{ svc}
}

func (u *customerCtrl) GetAll(ctx *gin.Context) {
	customers, err := u.svc.GetAll()
	if len(customers) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

func (u *customerCtrl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	customer, err := u.svc.GetByID(id)
	if customer == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func (u *customerCtrl) Store(ctx *gin.Context) {
	var customer customer.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.svc.Store(&customer)
	ctx.Status(http.StatusCreated)
}

func (u *customerCtrl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var customer customer.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer.ID = id
	u.svc.Update(&customer)
	ctx.Status(http.StatusOK)
}

func (u *customerCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	u.svc.Delete(id)
	ctx.Status(http.StatusNoContent)
}