package model

import (
	"github.com/pkg/errors"

	"github.com/YoriDigitalent/ImplementasiMVC-GO-24/app/utils"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type AuthModel struct {
	DB       *gorm.DB
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (auth AuthModel) Login() (bool, error, string) {
	var account AccountModel

	result := auth.DB.Where(&AccountModel{Name: auth.Name}).First(&account)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.Errorf("Account not found"), ""
		}

		return false, result.Error, ""
	}

	err := utils.HashComparator([]byte(account.Password), []byte(auth.Password))

	if err != nil {
		return false, errors.Errorf("incorrect password"), ""
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":           auth.Name,
		"account_number": account.AccountNumber,
	})

	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		return false, err, ""
	}

	return true, nil, token
}
