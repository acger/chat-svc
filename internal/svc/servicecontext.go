package svc

import (
	"github.com/acger/chat-svc/database"
	"github.com/acger/chat-svc/internal/config"
	"github.com/acger/user-svc/user"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	UserSvc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      database.NewMysql(&c),
		UserSvc: user.NewUser(zrpc.MustNewClient(c.UserSvc)),
	}
}
