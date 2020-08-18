package db

import (
	"fmt"
	orm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
	model "com.webfrm.app/models"
)

var dbob *DB

type DB struct {
	*model.DatabaseConfigurations
}

func NewDB(configurations *model.DatabaseConfigurations) {
	if dbob == nil {
		// read encrypted user name and password from config file
		// Decrypt it and pass it in the below string


		dbob = &DB{
			configurations,
		}
	}
}

func (db *DB) init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)


	// register model
	//orm.RegisterModel(new(User))

	// set default database

	err := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", db.DBUser, db.DBPassword, db.DBName), 30)

	if err != nil {

	}

	orm.RegisterModel(new(model.Patient))
}
