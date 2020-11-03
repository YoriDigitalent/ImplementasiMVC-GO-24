package model

import (
	"time"

	"github.com/YoriDigitalent/ImplementasiMVC-GO-24/app/constant"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type TransactionModel struct {
	DB                     *gorm.DB
	ID                     int    `gorm:"primary_key" json:"-"`
	TransactionType        int    `json:"transaction_type,omitempty"`
	TransactionDescription string `json:"transaction_description"`
	Sender                 int    `json:"sender"`
	Amount                 int    `json:"amount"`
	Recipient              int    `json:"recipient"`
	Timestamp              int64  `json:"timestamp"`
}

func (trx TransactionModel) Transfer() (bool, error) {

	err := trx.DB.Transaction(func(tx *gorm.DB) error {
		var sender, recipient AccountModel

		result := tx.Model(&AccountModel{}).Where(&AccountModel{
			AccountNumber: trx.Sender,
		}).First(&sender)

		if result.Error != nil {
			return result.Error
		}

		//Check balance first
		if sender.Saldo < trx.Amount {
			return errors.Errorf("Insufficient Saldo")
		}

		result.Update("saldo", sender.Saldo-trx.Amount)

		if result.Error != nil {
			return result.Error
		}

		result = tx.Model(&AccountModel{}).Where(AccountModel{
			AccountNumber: trx.Recipient,
		}).First(&recipient).
			Update("saldo", recipient.Saldo+trx.Amount)

		if result.Error != nil {
			return result.Error
		}

		trx.TransactionType = constant.TRANSFER
		trx.Timestamp = time.Now().Unix()

		result = tx.Create(&trx)
		if result.Error != nil {
			return result.Error
		}

		return nil

	})

	if err != nil {
		return false, err
	}

	return true, nil

}

func (trx TransactionModel) Withdraw() (bool, error) {

	err := trx.DB.Transaction(func(tx *gorm.DB) error {
		var sender AccountModel

		result := tx.Model(&AccountModel{}).Where(&AccountModel{
			AccountNumber: trx.Sender,
		}).First(&sender)

		if result.Error != nil {
			return result.Error
		}

		// Check balance first
		if sender.Saldo < trx.Amount {
			return errors.Errorf("Insufficient saldo")
		}

		result = result.Update("saldo", sender.Saldo-trx.Amount)
		if result.Error != nil {
			return result.Error
		}

		trx.TransactionType = constant.WITHDRAW
		trx.Timestamp = time.Now().Unix()
		result = tx.Create(&trx)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (trx TransactionModel) Deposit() (bool, error) {
	err := trx.DB.Transaction(func(tx *gorm.DB) error {
		var sender AccountModel
		result := tx.Model(&AccountModel{}).Where(&AccountModel{
			AccountNumber: trx.Sender,
		}).First(&sender).
			Update("saldo", sender.Saldo+trx.Amount)

		if result.Error != nil {
			return result.Error
		}

		trx.TransactionType = constant.DEPOSIT
		trx.Timestamp = time.Now().Unix()
		result = tx.Create(&trx)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
