package database

import (
	"github.com/acger/chat-svc/model"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(model.Chat{})
	db.AutoMigrate(model.ChatHistory{})
}
