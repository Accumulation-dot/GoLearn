package models

import "github.com/jinzhu/gorm"

type User struct {

	gorm.Model

	Name string

	Age uint8

	Sex uint
}
