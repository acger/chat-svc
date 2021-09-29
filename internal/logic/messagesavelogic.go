package logic

import (
	"context"
	"github.com/acger/chat-svc/model"

	"github.com/acger/chat-svc/internal/svc"
	"github.com/acger/chat-svc/template"

	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"
)

type MessageSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageSaveLogic {
	return &MessageSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageSaveLogic) MessageSave(in *template.MsgSaveReq) (*template.Rsp, error) {
	msg := model.Chat{}
	copier.Copy(&msg, in)
	l.svcCtx.DB.Create(&msg)

	return &template.Rsp{Code: 0}, nil
}
