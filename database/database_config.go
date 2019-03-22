package database

import "github.com/jinzhu/gorm"

func MysqlConfig()  {

	db, err := gorm.Open("mysql", "developer:ZXcV1357@$^*@/golang_gorm?charaset=utf8md4&parseTime=True&loc=Local")

	if err != nil {

	}

	defer db.Close()
}
