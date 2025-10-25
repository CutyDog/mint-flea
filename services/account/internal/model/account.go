package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UID string `gorm:"unique;not null"`
}
