package logic

import (
	"context"
	"github.com/acger/chat-svc/chat"
	"github.com/acger/chat-svc/internal/svc"
	"github.com/acger/chat-svc/model"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type MessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageListLogic {
	return &MessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageListLogic) MessageList(in *chat.MsgListReq) (*chat.MsgListRsp, error) {
	var result *chat.MsgListRsp

	l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		//更新已读状态
		tx.Model(&model.Chat{}).Where("uid = ?", in.ToUid).Where("to_uid = ?", in.Uid).Update("status", true)

		var list []*model.Chat

		tx = tx.Where(
			tx.Where("uid = ?", in.Uid).Where("to_uid = ?", in.ToUid),
		).Or(
			tx.Where("uid = ?", in.ToUid).Where("to_uid = ?", in.Uid),
		)

		/* 聊天记录分页尚有问题，暂时屏蔽
		if in.PageSize == 0 {
			in.PageSize = 30
		}

		if in.Page == 0 {
			in.Page = 1
		}

		offset := (in.Page - 1) * in.PageSize
		tx.Limit(int(in.PageSize)).Offset(int(offset)).Order("id desc").Find(&list)
		*/

		tx.Order("id").Find(&list)

		var total int64
		tx.Model(model.Chat{}).Count(&total)

		result = &chat.MsgListRsp{
			Code:     0,
			Total:    total,
			PageSize: in.PageSize,
			Page:     in.Page,
			Msg:      make([]*chat.ChatMessage, len(list)),
		}

		for i, item := range list {
			tmp := chat.ChatMessage{}
			copier.Copy(&tmp, &item)
			result.Msg[i] = &tmp
		}

		return nil
	})

	return result, nil
}
