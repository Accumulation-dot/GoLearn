package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func MysqlConfig() (db *gorm.DB, err error)  {
	db, err = gorm.Open("mysql", "developer:ZXcV1357@$^*@/golang_gorm?charaset=utf8md4&parseTime=True&loc=Local")
	return
}

func PostgreSQLCOnfig() (db *gorm.DB, err error)  {
	db, err = gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")

	return
}