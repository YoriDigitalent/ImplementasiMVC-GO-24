package model

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/YoriDigitalent/ImplementasiMVC-GO-24/app/utils"
	"gorm.io/gorm"
)

type Account struct {
	ID            int    `gorm:"primary_key" json:"-"`
	IdAccount     string `json:"id_account,omitempty"`
	Name          string `json:"name"`
	Password      string `json:"password,omitempty"`
	AccountNumber int    `json:"account_number,omitempty"`
	Saldo         int    `json:"saldo"`
}

type AccountModel struct {
	DB *gorm.DB
}

func (model AccountModel) InsertNewAccount(account Account) (bool, error) {
	account.AccountNumber = utils.RangeIn(1000, 999999)
	account.Saldo = 0
	account.IdAccount = fmt.Sprintf("id-%d", utils.RangeIn(10, 5000))

	result := model.DB.Create(&account)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (model AccountModel) GetAccountDetail(idAccount int) (bool, error, []TransactionModel, Account) {
	var transaction []TransactionModel
	var account Account

	result := model.DB.Model(&Transaction{}).Where("sender = ? OR recipient = ?", idAccount, idAccount).Find(&transaction)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.Errorf("Account not found"), []TransactionModel{}, Account{}
		}
		return false, result.Error, []TransactionModel{}, Account{}
	}

	result = model.DB.Where(&Account{
		AccountNumber: idAccount,
	}).Find(&AccountModel{})

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.Errorf("Account not found"), []TransactionModel{}, Account{}
		}
		return false, result.Error, []TransactionModel{}, Account{}
	}

	return true, nil, transaction, account

}
