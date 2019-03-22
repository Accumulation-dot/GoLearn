package models

import "github.com/jinzhu/gorm"

const (
	MenuCodeLast = "history"
	MenuCodeSlot = "slot"
	MenuCodeCheese = "cheese"
	MenuCodeReal = "real"
	MenuCodeFishing = "fishing"
	MenuCodeSports = "sports"
	MenuCodeLottery = "lottery"
)

type HMCode string

type HomeMenu struct {
	gorm.Model

	Title string `json:"title"`

	Img string `json:"img"`

	Status int8 `json:"status"`

	Code HMCode `json:"code"`
}