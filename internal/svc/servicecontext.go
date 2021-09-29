package svc

import (
	"github.com/acger/chat-svc/database"
	"github.com/acger/chat-svc/internal/config"
	"github.com/acger/user-svc/userclient"
	"github.com/tal-tech/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	UserSvc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      database.NewMysql(&c),
		UserSvc: userclient.NewUser(zrpc.MustNewClient(c.UserSvc)),
	}
}
