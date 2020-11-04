package controller

import (
	"net/http"

	"github.com/YoriDigitalent/ImplementasiMVC-GO-24/app/model"
	"github.com/YoriDigitalent/ImplementasiMVC-GO-24/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//TransactionController is struct to define database
type TransactionController struct {
	DB *gorm.DB
}

//Transfer is function to transfer
func (c TransactionController) Transfer(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: c.DB,
	}

	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Transfer(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}

//Withdraw is function to transfer
func (c TransactionController) Withdraw(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: c.DB,
	}

	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Withdraw(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}

//Deposit is function to transfer
func (c TransactionController) Deposit(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: c.DB,
	}

	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Deposit(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}
