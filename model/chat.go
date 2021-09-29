package model

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Uid     uint64 `gorm:"index"`
	ToUid   uint64 `gorm:"index"`
	Message string `gorm:"size:3000"`
	Status  bool
}

type ChatHistory struct {
	gorm.Model
	Uid   uint64 `gorm:"index"`
	ToUid uint64
}
