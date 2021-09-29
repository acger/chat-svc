package logic

import (
	"context"
	"github.com/acger/chat-svc/model"
	"gorm.io/gorm"

	"github.com/acger/chat-svc/internal/svc"
	"github.com/acger/chat-svc/template"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChatHistorySaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatHistorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistorySaveLogic {
	return &ChatHistorySaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatHistorySaveLogic) ChatHistorySave(in *template.CHSaveReq) (*template.Rsp, error) {
	db := l.svcCtx.DB
	db.Transaction(func(tx *gorm.DB) error {
		var fromHistory model.ChatHistory
		tx.Model(&model.ChatHistory{}).Where("uid = ?", in.Uid).Where("to_uid = ?", in.ToUid).Find(&fromHistory)
		if fromHistory.ID == 0 {
			tx.Model(&model.ChatHistory{}).Create(&model.ChatHistory{
				Uid:   in.Uid,
				ToUid: in.ToUid,
			})
		}

		var toHistory model.ChatHistory
		tx.Model(&model.ChatHistory{}).Where("uid = ?", in.ToUid).Where("to_uid = ?", in.Uid).Find(&toHistory)
		if toHistory.ID == 0 {
			tx.Model(&model.ChatHistory{}).Create(&model.ChatHistory{
				Uid:   in.ToUid,
				ToUid: in.Uid,
			})
		}

		return nil
	})

	return &template.Rsp{}, nil
}
