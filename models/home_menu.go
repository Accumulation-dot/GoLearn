package models

import "github.com/jinzhu/gorm"

const (
	HMCodeInfo = "info"



)

type HMCode string

type HomeMenu struct {
	gorm.Model

	Title string `json:"title"`

	Img string `json:"img"`

	Status int8 `json:"status"`

	Code HMCode `json:"code"`
}