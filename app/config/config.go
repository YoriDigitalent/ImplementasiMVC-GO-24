package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/YoriDigitalent/ImplementasiMVC-GO-24/app/model"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@/simple_bank?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB" + err.Error())
	}

	//automigrate here
	db.AutoMigrate(new(model.AccountModel), new(model.TransactionModel))

	return db
}