package logic

import (
	"context"
	"github.com/acger/chat-svc/model"

	"github.com/acger/chat-svc/internal/svc"
	"github.com/acger/chat-svc/chat"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatNumberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatNumberLogic {
	return &ChatNumberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatNumberLogic) ChatNumber(in *chat.ChatNumberReq) (*chat.ChatNumberRsp, error) {
	var num int64
	tx := l.svcCtx.DB.Model(&model.ChatHistory{})
	tx.Where("uid = ?", in.Id).Count(&num)

	return &chat.ChatNumberRsp{Code: 0, Number: num}, nil
}
