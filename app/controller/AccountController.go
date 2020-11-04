package controller

import (
	"log"
	"net/http"

	"github.com/YoriDigitalent/ImplementasiMVC-GO-24/app/model"
	"github.com/YoriDigitalent/ImplementasiMVC-GO-24/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//AccountController is struct to define database
type AccountController struct {
	DB *gorm.DB
}

//CreateAccount is function to create account
func (c AccountController) CreateAccount(ctx *gin.Context) {

	accountModel := model.AccountModel{
		DB: c.DB,
	}

	var account model.Account

	err := ctx.Bind(&account)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	hashPassword, err := utils.HashGenerator(account.Password)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	account.Password = hashPassword

	flag, err := accountModel.InsertNewAccount(account)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {

		utils.WrapAPIError(ctx, "Unknown failed to insert account", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
}

//GetAccount is function to get all account
func (c AccountController) GetAccount(ctx *gin.Context) {

	idAccount := ctx.MustGet("account_number").(int)
	accountModel := model.AccountModel{
		DB: c.DB,
	}

	flag, err, transactions, account := accountModel.GetAccountDetail(idAccount)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(ctx, map[string]interface{}{
		"account":     account,
		"transaction": transactions,
	}, http.StatusOK, "success")
	return
}

//Login is function to login in application
func (c AccountController) Login(ctx *gin.Context) {
	authModel := model.AuthModel{
		DB: c.DB,
	}

	var auth model.Auth

	err := ctx.Bind(&auth)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err, token := authModel.Login(auth)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(ctx, map[string]interface{}{
		"token": token,
	}, http.StatusOK, "success")

	log.Println("Login")

}
