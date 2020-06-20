package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name		string	`json:"name" validate:"required"`
	Phone		string	`json:"phone" validate:"required" gorm:"type:varchar(100);unique_index"	`
	Password	string	`json:"password"`
	Role		string 	`json:"role" validate:"required"`
}